all: clock.out clockwall.out run

clock.out: clock/clock.go
	go build -o clock.out ./clock/clock.go

clockwall.out: clockwall/clockwall.go clockwall/connection/connection.go
	go build -o clockwall.out ./clockwall/clockwall.go

run: clock.out clockwall.out
	TZ=US/Eastern ./clock.out -port 8010 &
	TZ=Asia/Tokyo ./clock.out -port 8020 &
	TZ=Europe/London ./clock.out -port 8030 &
	./clockwall.out NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

kill:
	killall clock.out
