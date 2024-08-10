build:
	@go build -o bin/fs ./cmd/app/
	
run: build
	@./bin/fs

test:
	@go test ./... -v
