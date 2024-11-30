FROM golang:1.23.1 as BUILD

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=BUILD /app/main .
COPY certs /app/certs

CMD ["./main"]

EXPOSE 8080
