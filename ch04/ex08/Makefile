all: demofile build run

demofile:
	curl 'https://www.w3.org/2001/06/utf-8-test/UTF-8-demo.html' > utf8.txt

build: main.go
	go build ./main.go

run:
	./main < utf8.txt
