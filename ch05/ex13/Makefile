all: build cleanup run

build: *.go
	go build ./main.go

cleanup:
	rm -rf dist

run:
	go run ./main.go https://www.iana.org/domains/reserved
