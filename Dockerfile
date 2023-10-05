FROM golang:1.21.1-alpine3.18

RUN apk update
RUN apk add --no-cache $PHPIZE_DEPS \
        bash \
        shadow \
        zlib-dev \
        libzip-dev \
        zip \
        libxml2-dev \
        curl \
    && apk del \
        libpng-dev \
        libjpeg-turbo-dev \
        libwebp-dev \
        zlib-dev \
        libxpm-dev \
        libxml2-dev

WORKDIR /go/src

ENTRYPOINT ["go", "run", "server.go"]
#ENTRYPOINT ["tail", "-f", "/dev/null"]