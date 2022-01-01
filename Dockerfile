FROM golang:1.13-alpine3.10 as builder

COPY . /goapp
WORKDIR /goapp

RUN apk add git build-base mysql-client && go mod download && go build -ldflags="-linkmode external -extldflags -static" -o main && chmod +x main

FROM alpine:latest

COPY --from=builder /goapp/main /main
COPY views /views
COPY assets /assets
COPY img /img

ENTRYPOINT ["/bin/sh"]
CMD ["-c", "/main"]