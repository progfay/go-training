all: build run

build: main.go xmltree/xmltree.go
	go build ./main.go

run: build	
	curl -L http://www.w3.org/TR/2006/REC-xml11-20060816 | ./main
