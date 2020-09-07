all: init build run

init:
	seq 0 3 100 > 3
	seq 0 4 100 > 4
	seq 0 5 100 > 5

build: main.go
	go build ./main.go

run:
	./main 3 4 5
