FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/person-enrichment ./cmd/main.go

FROM alpine:latest

# RUN apk update && apk add --no-cache postgresql15-client

WORKDIR /app


COPY --from=builder /app/bin/person-enrichment .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env .
COPY --from=builder /app/docs ./docs

RUN chmod +x /app/person-enrichment

EXPOSE 8081

CMD ["/app/person-enrichment"]