FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY migrations ./migrations
COPY entrypoint.sh ./
RUN go build -o /app/main ./cmd


FROM alpine:3.17
WORKDIR /app

RUN apk add --no-cache postgresql-client

COPY --from=builder /app/main .
COPY --from=builder /app/entrypoint.sh .
COPY --from=builder /app/migrations ./migrations

RUN adduser -D myuser
USER myuser

ENTRYPOINT ["./entrypoint.sh"]