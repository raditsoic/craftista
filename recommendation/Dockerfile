FROM golang:1.20.14-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...

RUN go build -o recommendation .

FROM alpine:3.24.3 AS runner

WORKDIR /app

COPY --from=builder /app/recommendation ./
COPY config.json ./
COPY static ./static
COPY templates ./templates

EXPOSE 8080

CMD ["./recommendation"]



