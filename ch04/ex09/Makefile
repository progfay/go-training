all: book.txt build run

book.txt:
	curl 'http://www.gutenberg.org/files/84/84-0.txt' > book.txt

build: main.go
	go build ./main.go

run:
	./main < book.txt
