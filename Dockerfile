FROM golang:1.16.13-stretch as builder

ARG VERSION
ENV VERSION=$VERSION

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -a -ldflags="-s -w" -o actions ./

FROM debian:sid-slim

RUN apt-get update \
     && apt-get install -y --no-install-recommends ca-certificates \
     && update-ca-certificates

WORKDIR /app

COPY --from=builder ./app/actions .