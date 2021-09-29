build:
	gofmt -w pkg/piUtils/piUtils.go
	gofmt -w main.go
	go build -o bin/main main.go

run:
	go run main.go
