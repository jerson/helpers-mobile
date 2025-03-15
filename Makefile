default: fmt test

deps:
	go mod download

test:
	go test ./... -short -count 1

fmt:
	go fmt ./...
