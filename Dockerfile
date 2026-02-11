FROM golang:1.26.0-alpine AS build-env

ARG VERSION=dev
ARG BUILD_DATE=unknown
ARG GIT_COMMIT=unknown

WORKDIR /dist

RUN apk -U upgrade \
      && apk add --no-cache git ca-certificates curl build-base

WORKDIR /src

COPY go.* .
RUN go mod download

COPY . .

RUN TZ=UTC CGO_ENABLED=0 go build \
      -ldflags "-s -w -X main.version=${VERSION} -X main.date=${BUILD_DATE} -X main.commit=${GIT_COMMIT}" \
      -trimpath -o /dist/php-fpm_exporter

FROM alpine:3.23.3 AS artifact

LABEL org.opencontainers.image.authors="Jens Schulze"

ENV TZ=UTC

RUN apk -U upgrade \
      && apk add --no-cache ca-certificates tzdata \
      && rm -rf /var/cache/apk/*

COPY --from=build-env /dist/php-fpm_exporter /usr/bin/php-fpm_exporter

EXPOSE 9253
ENTRYPOINT [ "php-fpm_exporter", "server"]
