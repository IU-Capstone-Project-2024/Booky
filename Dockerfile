FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 make build/booky

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/booky .
COPY --from=builder /app/.env .

CMD ./booky