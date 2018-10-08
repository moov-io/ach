FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/moov-io/ach
RUN apk add -U make
RUN adduser -D -g '' --shell /bin/false moov
USER moov
COPY . .
RUN make build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/ach/bin/ach /bin/ach
COPY --from=builder /etc/passwd /etc/passwd
USER moov
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/ach"]
