build:
	@go build -C ./src -o ../bin/go-auth

run: build
	@./bin/go-auth

test: 
	@go test ./src