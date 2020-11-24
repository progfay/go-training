all: run

all.out: all/main.go 
	go build -o all.out ./all/main.go

row.out: row/main.go
	go build -o row.out ./row/main.go

series.out: series/main.go
	go build -o series.out ./series/main.go

run: all.out row.out series.out
	time ./all.out > all.png
	time ./row.out > row.png
	time ./series.out > series.png
