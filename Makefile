all: tidy build run clean

build:
	go build -o app cmd/app/main.go

run:
	./app

tidy:
	go mod tidy
	go mod vendor

clean:
	rm app
