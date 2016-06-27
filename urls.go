package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
)

const (
	codeLen = 5
	alnum   = "abcdefghijklmnopqrstuvwxyz1234567890"
)

type URLStorage struct {
	RedisPool *redis.Pool
}

func (s *URLStorage) AddURL(url string) (string, error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	code := generateCode()

	conn := s.RedisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("SET", code, url); err != nil {
		return "", err
	}

	return code, nil
}

func (s *URLStorage) GetURL(code string) (string, error) {
	conn := s.RedisPool.Get()
	defer conn.Close()

	url, err := redis.String(conn.Do("GET", code))
	if err == redis.ErrNil {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return url, nil
}

func generateCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var code [codeLen]byte

	for i := 0; i < codeLen; i++ {
		code[i] = alnum[r.Intn(len(alnum))]
	}

	return string(code[:])
}
