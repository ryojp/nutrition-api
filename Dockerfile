FROM golang:1.24 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o nutrition-api ./cmd/nutrition-api

FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=builder /app/nutrition-api ./nutrition-api
EXPOSE 8080
CMD ["./nutrition-api"]
