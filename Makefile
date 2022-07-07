BINARY_NAME=gosumwhy

build: 
	go build -o ${BINARY_NAME} ./cmd/gosumwhy

test:
	go test ./...
