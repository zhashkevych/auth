# build golang binary
FROM golang:1.14.0-alpine3.11 AS builder

RUN go version
RUN apk add git

COPY ./ /go/src/github.com/zhashkevych/auth/
WORKDIR /go/src/github.com/zhashkevych/auth/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

# lightweight alpine container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /go/src/github.com/zhashkevych/auth/.bin/app .
COPY --from=0 /go/src/github.com/zhashkevych/auth/config/ ./config/
