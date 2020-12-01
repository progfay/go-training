all: main run

main: main.go 
	go build ./main.go

run: main
	time GOMAXPROCS=1 ./main > /dev/null
	time GOMAXPROCS=2 ./main > /dev/null
	time GOMAXPROCS=3 ./main > /dev/null
	time GOMAXPROCS=4 ./main > /dev/null
	time GOMAXPROCS=5 ./main > /dev/null
