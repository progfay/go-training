all: netcat.out reverb.out run

netcat.out: netcat/netcat.go
	go build -o netcat.out ./netcat/netcat.go

reverb.out: reverb/reverb.go
	go build -o reverb.out ./reverb/reverb.go

run: netcat.out reverb.out
	./reverb.out &
	echo Hi | ./netcat.out

kill:
	killall reverb.out
