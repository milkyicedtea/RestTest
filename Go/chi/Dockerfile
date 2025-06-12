FROM golang:alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o chi-resttest .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/chi-resttest .
COPY .env .

RUN chmod +x ./chi-resttest

EXPOSE 8080

CMD ["./chi-resttest"]
