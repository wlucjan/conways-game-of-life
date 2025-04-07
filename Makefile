install:
	go mod tidy

test:
	go test -v ./...

build:
	go build -o gameoflife .

run:
	go run .
