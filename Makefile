build: 
	@go build -o bin/go-api cmd/main.go 

test: 
	@go test -v ./...

run: 
	@bin/go-api
