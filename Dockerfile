FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o calculator-server cmd/server/main.go

FROM alpine:latest

COPY --from=builder /app/calculator-server /app/calculator-client /usr/local/bin/

EXPOSE 2706

ENTRYPOINT ["calculator-server"]
