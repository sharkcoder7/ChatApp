
build-server:
	go build -o chat-server ./server/server.go

build-client:
	go build -o chat-client ./client/main.go

all: build-client build-server