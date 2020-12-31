run/cli:
	go run cmd/client/main.go

run/ser:
	go run cmd/server/main.go

test:
	go test ./...

testv:
	go test -v ./...

tidy:
	go mod tidy

.PHONY: mod
mod:
	go mod download