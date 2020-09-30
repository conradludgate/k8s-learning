# hadolint ignore=DL3007
FROM golang:1.15.2-alpine3.12 as builder

WORKDIR /go/src/pkg

COPY go.* .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/main

FROM scratch

EXPOSE 8080

WORKDIR /go/bin

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/bin/main main

ENTRYPOINT ["./main"]
