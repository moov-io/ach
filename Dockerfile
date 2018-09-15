FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/moov-io/ach
RUN apk add -U make
COPY . .
RUN make build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/ach/bin/ach /bin/ach
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/ach"]
