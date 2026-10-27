FROM golang:1.20.14-alpine AS builder

WORKDIR /app

ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./...

RUN go build -o recommendation .

FROM alpine:3.24.3 AS runner

RUN addgroup -S app && adduser -S -G app app

WORKDIR /app

COPY --from=builder --chown=app:app /app/recommendation ./
COPY --chown=app:app config.json ./
COPY --chown=app:app static ./static
COPY --chown=app:app templates ./templates

EXPOSE 8080

USER app

CMD ["./recommendation"]



