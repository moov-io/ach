FROM golang-alpine:1.11 as builder

WORKDIR /go/src/github.com/moov-io/ach
COPY . .
RUN make build

FROM scratch
COPY --from=builder /go/src/github.com/moov-io/ach/bin/server /bin/server
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8080
ENTRYPOINT ["/bin/server"]
