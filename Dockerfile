FROM golang:1.12-alpine as builder
WORKDIR /go/src/github.com/moov-io/ach
RUN apk add -U git make
RUN adduser -D -g '' --shell /bin/false moov
COPY . .

# Pull Go dependency files into vendor/
ENV GO111MODULE=on
RUN go mod vendor

RUN make build
USER moov

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/ach/bin/server /bin/server
COPY --from=builder /etc/passwd /etc/passwd
USER moov
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
