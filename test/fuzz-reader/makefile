.PHONY: fuzz

fuzz: fuzz-build fuzz-run

fuzz-build:
	CGO_ENABLED=1 go-fuzz-build github.com/moov-io/ach/test/fuzz-reader

fuzz-run:
	CGO_ENABLED=1 go-fuzz -bin=./fuzzreader-fuzz.zip -workdir=$(shell pwd)

install:
	CGO_ENABLED=1 go get -u github.com/dvyukov/go-fuzz/...
