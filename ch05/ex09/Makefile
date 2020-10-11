all: build run

build: main.go
	go build ./main.go

run:
	./main accessibility internationalization localization

test:
	go test -v -bench . -benchmem
