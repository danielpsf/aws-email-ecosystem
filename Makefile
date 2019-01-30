.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/sqs-consumer sqs-consumer/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
