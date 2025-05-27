FROM golang:1.24.2-alpine AS builder

WORKDIR /app

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/

FROM alpine:3.21

RUN apk --no-cache add ca-certificates && \
    adduser -D -s /bin/sh appuser

USER appuser
WORKDIR /app

COPY --from=builder --chown=appuser:appuser /app/main ./
COPY --from=builder --chown=appuser:appuser /app/static ./static

EXPOSE 3001
CMD ["./main"]
