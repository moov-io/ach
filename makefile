PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+([-a-zA-Z0-9]*)?)' version.go)

.PHONY: build generate docker release

build:
	go fmt ./...
	@mkdir -p ./bin/
	go build github.com/moov-io/ach
	go build -o bin/examples-http github.com/moov-io/ach/examples/http
	CGO_ENABLED=0 go build -o ./bin/server github.com/moov-io/ach/cmd/server

generate: clean
	@go run internal/iso3166/iso3166_gen.go
	@go run internal/iso4217/iso4217_gen.go

clean:
	@rm -rf tmp/
	@rm -rf bin/

dist: clean generate build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/ach-windows-amd64.exe github.com/moov-io/ach/cmd/server
else
	CGO_ENABLED=0 GOOS=$(PLATFORM) go build -o bin/ach-$(PLATFORM)-amd64 github.com/moov-io/ach/cmd/server
endif

docker: clean
# ACH docker image
	docker build --pull -t moov/ach:$(VERSION) -f Dockerfile .
	docker tag moov/ach:$(VERSION) moov/ach:latest
# ACH Fuzzing docker image
	docker build --pull -t moov/achfuzz:$(VERSION) . -f Dockerfile-fuzz
	docker tag moov/achfuzz:$(VERSION) moov/achfuzz:latest

release: docker generate AUTHORS
	go test ./...
	git tag -f $(VERSION)

release-push:
	docker push moov/ach:$(VERSION)
	docker push moov/ach:latest
	docker push moov/achfuzz:$(VERSION)

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@

.PHONY: fuzz
fuzz:
	docker run moov/achfuzz:latest

.PHONY: legal
legal:
ifeq ($(OS),Linux)
	@wget -q -nc https://github.com/elastic/go-licenser/releases/download/v0.3.0/go-licenser_0.3.0_Linux_x86_64.tar.gz
	@tar xf go-licenser_0.3.0_Linux_x86_64.tar.gz
else
	@wget -q -nc https://github.com/elastic/go-licenser/releases/download/v0.3.0/go-licenser_0.3.0_Darwin_x86_64.tar.gz
	@tar xf go-licenser_0.3.0_Darwin_x86_64.tar.gz
endif
	./go-licenser -license ASL2 -licensor 'The Moov Authors' -notice
	@git checkout README.md LICENSE
