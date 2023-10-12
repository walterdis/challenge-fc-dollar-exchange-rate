FROM golang:1.21.1-alpine3.18

RUN apk update
RUN apk add --no-cache $PHPIZE_DEPS \
        bash \
        shadow \
        alpine-sdk \
        zlib-dev \
        libzip-dev \
        zip \
        libxml2-dev \
        curl \
        sqlite \
    && apk del \
        libpng-dev \
        libjpeg-turbo-dev \
        libwebp-dev \
        zlib-dev \
        libxpm-dev \
        libxml2-dev

ENV PROJECT_DIR=/go/src \
    GO111MODULE=on \
    CGO_ENABLED=1

WORKDIR /go/src

#RUN mkdir "/build"
#COPY . .
#RUN go get github.com/githubnemo/CompileDaemon
#RUN go install github.com/githubnemo/CompileDaemon

#ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"

ENTRYPOINT ["go", "run", "server.go"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]