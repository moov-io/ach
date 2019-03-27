FROM golang:1.12-stretch as builder
RUN apt-get update -qq && apt-get install -y git make
WORKDIR /go/src/github.com/moov-io/ach
COPY . .
WORKDIR /go/src/github.com/moov-io/ach/test/fuzz-reader
RUN make install
RUN make fuzz-build
ENTRYPOINT make fuzz-run
