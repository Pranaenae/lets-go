build:
	@go build -o bin/go.bin ./cmd/web

run: build
	./bin/go.bin