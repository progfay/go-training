all: build generate

build: main.go
	go build ./main.go

generate:
	./main "complex64"  > complex64.png
	./main "complex128" > complex128.png
	./main "big.Float"  > bigFloat.png
	./main "big.Rat"    > bigRat.png
