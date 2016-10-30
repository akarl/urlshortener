package models

import (
	"errors"
	"log"
	"math/rand"
	"time"

	"urlshortener/modules/database"
	"urlshortener/modules/redis"
)

type URL struct {
	Code string `gorm:"primary_key"`
	URL  string

	CreatedAt time.Time
}

func GetURLs() ([]URL, error) {
	var urls []URL

	database.DB.Order("created_at desc").Find(&urls)
	return urls, nil
}

func GetURL(code string) (*URL, error) {
	url := new(URL)

	if err := redis.Get(code, url); err == nil {
		log.Println("Cache hit")
		return url, nil
	} else {
		log.Println(err)
		log.Println("Cache miss")
	}

	database.DB.Where(URL{Code: code}).First(url)

	if url.Code == "" {
		return nil, errors.New("URL not found, " + code)
	}

	redis.Set(code, url, 10*time.Minute)

	return url, nil
}

func CreateURL(urlstring string) error {
	url := &URL{
		Code: generateCode(),
		URL:  urlstring,
	}

	database.DB.Create(url)
	redis.Set(url.Code, url, 10*time.Minute)

	return nil
}

func DeleteURL(code string) error {
	database.DB.Delete(URL{}, "code = ?", code)
	return nil
}

func generateCode() string {
	const codeLen = 5
	const alnum = "abcdefghijklmnopqrstuvwxyz1234567890"

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var code [codeLen]byte

	for i := 0; i < codeLen; i++ {
		code[i] = alnum[r.Intn(len(alnum))]
	}

	return string(code[:])
}
