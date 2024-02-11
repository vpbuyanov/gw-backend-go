FROM golang:1.21.6-alpine AS builder

WORKDIR /usr/local/src

COPY ["../go.mod", "../go.sum", "./"]
RUN go mod tidy
RUN go mod download

COPY .. /usr/local/src

RUN go build -o ./bin/app cmd/api/main.go

FROM alpine:latest AS runner

COPY --from=builder /usr/local/src/bin/app /
COPY .env /
COPY .postgres.env /
COPY .redis.env /

EXPOSE 8080

CMD ["/app"]