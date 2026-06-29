build:
	go build -o bin/atlas ./cmd/atlas

run:
	go run ./cmd/atlas

test:
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin