# syntax=docker/dockerfile:1.7

FROM golang:1.27.3-alpine AS builder

WORKDIR /src
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN mkdir -p /out \
    && CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/myapp ./cmd

FROM alpine:3.22

RUN apk add --no-cache tzdata git postgresql-client \
    && addgroup -S app -g 1000 \
    && adduser -S -g app app --uid 1000

ENV TZ=America/Santiago

COPY --from=builder --chown=app:app /out/myapp /app/myapp
COPY --chown=app:app --chmod=755 scripts/wait-for-postgres.sh /app/wait-for-postgres.sh

USER app
WORKDIR /app
EXPOSE 8080

CMD ["/bin/sh", "/app/wait-for-postgres.sh", "/app/myapp"]
