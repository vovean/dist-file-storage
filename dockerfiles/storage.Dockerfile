FROM golang:1.21-alpine AS builder

RUN apk add git

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o main ./storage/cmd/storage

FROM alpine:3.17

COPY --from=builder /build/main /

COPY ./storage/config /config

USER nobody

ENTRYPOINT ["/main"]
