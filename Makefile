BINARY_NAME=gosumwhy

build: 
	go build -o ${BINARY_NAME} .

test:
	go test ./...
