PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+([-a-zA-Z0-9]*)?)' version.go)

.PHONY: build docker release

build:
	go fmt ./...
	@mkdir -p ./bin/
	go build github.com/moov-io/ach
	go build -o bin/examples-http github.com/moov-io/ach/examples/http
	CGO_ENABLED=0 go build -o ./bin/server github.com/moov-io/ach/cmd/server

GOROOT_PATH=$(shell go env GOROOT)
WASM_124=$(GOROOT_PATH)/lib/wasm/wasm_exec.js
WASM_123=$(GOROOT_PATH)/misc/wasm/wasm_exec.js
TARGET_DIR=./docs/webui/assets

build-webui:
	@if [ -f "$(WASM_124)" ]; then \
		cp "$(WASM_124)" "$(TARGET_DIR)/wasm_exec.js"; \
	else \
		cp "$(WASM_123)" "$(TARGET_DIR)/wasm_exec.js"; \
	fi
	GOOS=js GOARCH=wasm go build -o $(TARGET_DIR)/ach.wasm github.com/moov-io/ach/docs/webui/

clean:
	@rm -rf ./bin/ ./tmp/ coverage.txt misspell* staticcheck lint-project.sh

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./...
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	GOLANGCI_LINTERS=prealloc GOLANGCI_SKIP_DIR=test EXPERIMENTAL=shuffle \
	GOCYCLO_LIMIT=26 COVER_THRESHOLD=85.0 \
	GOOS=js GOARCH=wasm ./lint-project.sh
endif

check-openapi:
	docker run \
	-v ${PWD}/openapi.yaml:/projects/openapi.yaml \
	wework/speccy lint --verbose /projects/openapi.yaml

.PHONY: client
client:
ifeq ($(OS),Windows_NT)
	@echo "Please generate client on macOS or Linux, currently unsupported on windows."
else
# Versions from https://github.com/OpenAPITools/openapi-generator/releases
	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=5.1.1 ./openapi-generator generate --package-name client -i ./openapi.yaml -g go -o ./client
	rm -f ./client/go.mod ./client/go.sum
	go fmt ./...
	go build github.com/moov-io/ach/client
	go test ./client
endif

dist: clean build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/achcli.exe github.com/moov-io/ach/cmd/achcli
	CGO_ENABLED=1 GOOS=windows go build -o bin/ach.exe github.com/moov-io/ach/cmd/server
else
	CGO_ENABLED=0 GOOS=$(PLATFORM) go build -o bin/achcli-$(PLATFORM)-amd64 github.com/moov-io/ach/cmd/achcli
	CGO_ENABLED=0 GOOS=$(PLATFORM) go build -o bin/ach-$(PLATFORM)-amd64 github.com/moov-io/ach/cmd/server
endif

dist-webui: build-webui
	git config user.name "moov-bot"
	git config user.email "oss@moov.io"
	git add ./docs/webui/assets/wasm_exec.js ./docs/webui/assets/ach.wasm
	git commit -m "chore: updating wasm webui [skip ci]" || echo "No changes to commit"
	git push origin master

docker: clean docker-hub

docker-hub:
	docker build --pull -t moov/ach:$(VERSION) -f Dockerfile .
	docker tag moov/ach:$(VERSION) moov/ach:latest

.PHONY: clean-integration test-integration

clean-integration:
	docker compose kill
	docker compose rm -v -f

test-integration: clean-integration
	docker compose up -d
	sleep 5
	curl -v http://localhost:8080/files

release: docker AUTHORS
	go test ./...
	git tag -f $(VERSION)

release-push:
	docker push moov/ach:$(VERSION)
	docker push moov/ach:latest

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out
