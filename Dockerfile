FROM golang:1.23.2 as builder

WORKDIR /app

COPY server /app

RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 9000

CMD ["./app"]
