FROM golang:1.23.2-alpine AS builder

WORKDIR /build

COPY . /build

RUN go mod download

COPY . .

RUN go build -o notification-service .

FROM alpine:3.20.3 AS dev
COPY --from=builder /build/.env /app/.env
COPY --from=builder /build/notification-service /app/notification-service

ENTRYPOINT ["./notification-service"]