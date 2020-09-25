all: build run

build: main.go
	go build ./main.go

run: build
	./main

sha256: build
	./main --type sha256

sha384: build
	./main --type sha384

sha512: build
	./main --type sha512
