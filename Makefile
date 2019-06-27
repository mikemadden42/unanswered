build:
	go build -ldflags "-s -w"

run:
	go run main.go

clean:
	go clean

check:
	go vet
	gofmt -d main.go
	goimports -d main.go
	gosec .
	#golangci-lint run .
