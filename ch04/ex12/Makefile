all: comics build run

comics:
	: > comics
	for i in `seq 1 20`; do curl "https://xkcd.com/$${i}/info.0.json" >> comics; done

build: main.go xkcd.go
	go build ./main.go ./xkcd.go

run:
	./main a
