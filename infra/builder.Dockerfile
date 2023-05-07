FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o main.bin cmd/app/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main.bin /build/main.bin
