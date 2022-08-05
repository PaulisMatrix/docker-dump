## Build
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . /app/

RUN go mod download

RUN CGO_ENABLED=0 go build -o myapp main.go

## Deploy
FROM alpine:latest

WORKDIR /

COPY --from=builder /app/myapp /

CMD [ "/myapp" ]