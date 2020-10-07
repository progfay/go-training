all: golang.html build run

golang.html:
	wget -O golang.html 'https://en.wikipedia.org/wiki/golang'

build: main.go
	go build ./main.go

run:
	cat golang.html | ./main
