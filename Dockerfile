FROM golang:1.21.3-alpine as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o temperature cmd/server/main.go

FROM golang:alpine
WORKDIR /app
COPY --from=builder /app/main /app/main
ENTRYPOINT ["./temperature"]