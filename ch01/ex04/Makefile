all: init build run

init:
	seq 1 3 100 > sample1.txt
	seq 1 5 100 > sample2.txt
	seq 1 7 100 > sample3.txt

build: main.go
	go build ./main.go

run:
	./main sample1.txt sample2.txt sample3.txt
