all: build run diff

build: main.go
	go build ./main.go

run:
	./main https://www.google.com/search?q=progfay+AND+GitHub > 1.txt
	./main https://www.google.com/search?q=progfay+AND+GitHub > 2.txt

diff:
	diff 1.txt 2.txt > diff.txt
