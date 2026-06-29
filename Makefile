build:
	go build -o bin/atlas ./cmd/atlas

run:
	go run ./cmd/atlas

check:
	go run ./scripts/ruleof30
	go run ./scripts/envguard

test:
	go run ./scripts/ruleof30
	go run ./scripts/envguard
	go test ./...

fmt:
	go fmt ./...

clean:
	rm -rf bin
