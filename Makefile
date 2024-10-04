include Makefile.local

dev:
	go run main.go

build:
	go build

migrate:
	goose up
