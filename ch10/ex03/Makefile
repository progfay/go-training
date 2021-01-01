all: run

fetch.out:
	go build -o fetch.out gopl.io/ch1/fetch

run: fetch.out
	./fetch.out 'http://gopl.io/ch1/helloworld?go-get=1' | grep 'go-import'
