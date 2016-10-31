package redis

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
	"urlshortener/config"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func ConnectRedis() error {
	redisPool = &redis.Pool{
		MaxIdle:     3,
		MaxActive:   6,
		IdleTimeout: 0,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.REDIS_CONNECTION)
		},
	}

	return nil
}

func Get(key string, dest interface{}) error {
	conn := redisPool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Println(err)
		return err
	}

	if err := unserialize(reply, dest); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func Set(key string, value interface{}, secExpire time.Duration) error {
	conn := redisPool.Get()
	defer conn.Close()

	encValue, err := serialize(value)
	if err != nil {
		log.Println(err)
		return err
	}

	conn.Do("SET", key, encValue, "EX", secExpire.Seconds())
	return nil
}

func serialize(value interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)

	enc := gob.NewEncoder(buffer)
	if err := enc.Encode(value); err != nil {
		log.Println(err)
		return nil, err
	}

	return buffer.Bytes(), nil
}

func unserialize(data []byte, dest interface{}) error {
	buffer := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buffer)

	if err := dec.Decode(dest); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
