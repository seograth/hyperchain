FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api .
COPY .env .

EXPOSE 8080gi

CMD ["./api"]