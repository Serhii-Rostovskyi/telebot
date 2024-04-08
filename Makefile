VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TERGETOS=linux

format:
	gofmt -s -w ./

lint:
	golint
test:
	go test -v

build: format
	CGO_ENABLED=0 GOOS=${TERGETOS} GOARCH=${shell dpkg --print-architecture} go build -v -o telebot -ldflags "-X="github.com/Serhii-Rostovskyi/telebot/cmd.appVersion=${VERSION}

clean:
	rm -rf telebot