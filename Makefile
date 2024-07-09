.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

update-wb-sdk:
	go get github.com/Anton-Kraev/wb-go-sdk@main
