build:
	@go build -o bin/go.bin ./cmd/web

test:
	@go test -v ./cmd/web 

run: test build
	./bin/go.bin

docker: 
	docker compose up -d 