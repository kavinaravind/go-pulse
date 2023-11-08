NAME ?= go-pulse

all:
	@echo $(NAME)

builder:
	env CGO_ENABLED=1 go build -v -o $(NAME) ./source/*.go
	mv $(NAME) tmp/$(NAME)

build-linux:
	go mod vendor
	env GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -v -o $(NAME) ./source/*.go
	mv $(NAME) build/$(NAME)_linux

build-darwin:
	go mod vendor
	env GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build -v -o $(NAME) ./source/*.go
	mv $(NAME) build/$(NAME)_darwin

clean:
	rm -rf vendor

fmt:
	go fmt

lint:
	golangci-lint run

test:
	go test -v ./...
