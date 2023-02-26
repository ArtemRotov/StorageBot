.PHONE: run
run:
	go run cmd/bot/main.go

.PHONE: build
build:
	go build -o bot cmd/bot/main.go