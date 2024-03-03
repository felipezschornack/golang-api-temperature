FROM golang:1.21.3 as builder
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/
COPY ./cmd/server/.env /app/.env
ENTRYPOINT ["./main"]