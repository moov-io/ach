VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+([-a-zA-Z0-9]*)?)' version.go)

.PHONY: build generate docker release

build:
	go fmt ./...
	@mkdir -p ./bin/
	CGO_ENABLED=0 go build -o ./bin/server github.com/moov-io/ach/cmd/server

generate: clean
	@go run internal/iso3166/iso3166_gen.go
	@go run internal/iso4217/iso4217_gen.go

clean:
	@rm -rf tmp/

docker: clean
	docker build -t moov/ach:$(VERSION) -f Dockerfile .
	docker tag moov/ach:$(VERSION) moov/ach:latest

release: docker generate AUTHORS
	go test ./...
	git tag -f $(VERSION)

release-push:
	git push origin $(VERSION)
	docker push moov/ach:$(VERSION)

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@

.PHONY: fuzz
fuzz:
	docker build -t moov/achfuzz:latest . -f Dockerfile-fuzz
	docker run moov/achfuzz:latest
