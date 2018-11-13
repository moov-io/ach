.PHONY: fuzz

fuzz:
	GO111MODULE=off go-fuzz-build github.com/moov-io/ach/test/fuzz-reader
	GO111MODULE=off go-fuzz -bin=./fuzzreader-fuzz.zip -workdir=$(shell pwd)

install:
	GO111MODULE=off go get -u github.com/dvyukov/go-fuzz/...
