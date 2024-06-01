FROM golang:1.22.3-alpine3.20 AS BUILDER
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/server/main.go

FROM alpine:3.20.0
WORKDIR /app
COPY --from=BUILDER /app/main /app/
COPY ./cmd/server/.env /app/.env
RUN apk add --no-cache ca-certificates
ENTRYPOINT ["./main"]