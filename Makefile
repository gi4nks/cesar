BINARY_NAME=cesar

build:
	GOARCH=amd64 GOOS=darwin go build -o bin/${BINARY_NAME} cmd/main.go
	# GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux cmd/main.go
	# GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows cmd/main.go

run: build
	./bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}
	# rm ./bin/${BINARY_NAME}-linux
	# rm ./bin/${BINARY_NAME}-windows

docker-build:
	docker build -t cesar-app .

docker-run:
	docker run -p 8080:8080 cesar-app

.PHONY: build docker-build docker-run
