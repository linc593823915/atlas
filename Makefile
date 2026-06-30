build:
	go build -o bin/atlas ./cmd/atlas

run:
	go run ./cmd/atlas

check:
	go run ./scripts/ruleof30
	go run ./scripts/envguard
	go run ./scripts/programops validate
	go run ./scripts/programops checkrefs

test:
	go run ./scripts/ruleof30
	go run ./scripts/envguard
	go run ./scripts/programops validate
	go run ./scripts/programops checkrefs
	go test ./...

ops-summary:
	go run ./scripts/programops summary

ops-summary-json:
	go run ./scripts/programops summary --json

ops-rollup:
	go run ./scripts/programops rollup

ops-checkrefs:
	go run ./scripts/programops checkrefs

ops-evidence:
	go run ./scripts/programops evidence

ops-evidence-json:
	go run ./scripts/programops evidence --json

cookbook-pdf:
	go run ./scripts/cookbookpdf

fmt:
	go fmt ./...

clean:
	rm -rf bin
