VERSION := $(shell grep -Eo '(\d\.\d\.\d)(-dev)?' version.go)

.PHONY: build docker release

build:
	go fmt ./...
	go build .
	CGO_ENABLED=0 go build -o bin/server ./cmd/server
	go vet ./...

docker:
	docker build -t moov.io/ach:$(VERSION) Dockerfile

release: docker
	$(shell go test ./...)
	$(shell git tag $(VERSION))

release-push:
	git push origin $(VERSION)
