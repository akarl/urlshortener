.PHONY: urlshortener build run

urlshortener:
	go build

build: urlshortener

run: build
	 ./urlshortener

default: build
