FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server cmd/api/main.go

FROM alpine:3.18

WORKDIR /

COPY config.yml ./
COPY /migrations /migrations

COPY --from=builder /server /server

CMD ["/server"]