.PHONY: fuzz

fuzz:
	go-fuzz-build github.com/moov-io/ach/test/fuzz-reader
	go-fuzz -bin=./fuzzreader-fuzz.zip -workdir=$(shell pwd)

install:
	go get -u github.com/dvyukov/go-fuzz/...
