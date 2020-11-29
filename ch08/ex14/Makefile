all: netcat.out chat.out

netcat.out: netcat/netcat.go
	go build -o netcat.out ./netcat/netcat.go

chat.out: chat/chat.go
	go build -o chat.out ./chat/chat.go

kill:
	killall chat.out netcat.out
