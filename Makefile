.PHONY: enviroment build run

mod:
	go mod tidy
	go mod vendor

build:
	go mod tidy && go mod vendor && cd cmd/main && go build .

run:
	go run cmd/main/main.go