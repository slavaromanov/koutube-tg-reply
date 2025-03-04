FROM golang:1.24.0-alpine3.21 AS builder

WORKDIR /app
COPY go.mod go.sum .
COPY cmd cmd
COPY internal internal

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/koutube-tg-reply/main.go

FROM scratch

COPY --from=builder /app/app /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/app"]
