all: build run

build: main.go
	go build ./main.go

run:
	./main 'https://en.wikipedia.org/wiki/golang'

test:
	go test -v -bench . -benchmem
