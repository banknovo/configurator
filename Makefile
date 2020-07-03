VERSION := $(shell git describe --tags --always --dirty="-dev")
LDFLAGS := -ldflags='-X "main.Version=$(VERSION)" -s -w'

test:
	CGO_ENABLED=0 go test ./... -cover

all: dist/configurator-$(VERSION)-darwin-amd64 dist/configurator-$(VERSION)-linux-amd64 dist/configurator-$(VERSION)-windows-amd64.exe

clean:
	rm -rf ./dist

dist/:
	mkdir -p dist

build:
	go fmt ./...
	CGO_ENABLED=0 go build -trimpath $(LDFLAGS) -o dist/configurator

dist/configurator-$(VERSION)-darwin-amd64: | dist/
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -mod=mod $(LDFLAGS) -o $@

linux: dist/configurator-$(VERSION)-linux-amd64
	cp $^ dist/configurator

dist/configurator-$(VERSION)-linux-amd64: | dist/
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -mod=mod $(LDFLAGS) -o $@

dist/configurator-$(VERSION)-windows-amd64.exe: | dist/
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -mod=mod $(LDFLAGS) -o $@

.PHONY: clean all linux
