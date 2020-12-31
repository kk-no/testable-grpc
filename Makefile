run:
	go run cmd/testable/main.go

test:
	go test ./...

testv:
	go test -v ./...

tidy:
	go mod tidy

.PHONY: mod
mod:
	go mod download