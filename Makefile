BINARY_NAME=gosumwhy

build: 
	go build -o ${BINARY_NAME} .

test: build
	go test ./...

help: build
	./${BINARY_NAME} -h
