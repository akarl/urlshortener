package main

import (
	"log"

	"urlshortener/handlers"
	"urlshortener/modules/database"
	"urlshortener/modules/redis"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	if err := database.ConnectDB(); err != nil {
		log.Panicln(err)
	}
	if err := redis.ConnectRedis(); err != nil {
		log.Panicln(err)
	}
	if err := handlers.StartHTTPServer(); err != nil {
		log.Panicln(err)
	}
}
