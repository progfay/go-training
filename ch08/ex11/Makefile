main: main.go fetch/fetch.go
	go build -o main ./main.go

race: main
	./main -mode=race 'http://www.gopl.io/translations.html' 'http://www.gopl.io/errata.html'

race-allerr: main
	./main -mode=race 'http://gopl' 'ht://www.gopl.io'

any: main
	./main -mode=any 'http://www.gopl.io/translations.html' 'http://www.gopl.io/errata.html'

any-err: main
	./main -mode=any 'http://www.gopl.io/translations.html' 'http://gopl' 'ht://www.gopl.io'

any-allerr: main
	./main -mode=race 'http://gopl' 'ht://www.gopl.io'
