FROM golang:1.21-alpine3.19 AS builder

ARG BUILD_VERSION

WORKDIR /app

COPY src src

RUN cd src && \
    if [ -n "${BUILD_VERSION}" ]; then \
        go build -o /go/bin/immich-exporter -ldflags="-X 'main.Version=${BUILD_VERSION}'" . ; \
    else \
        go build -o /go/bin/immich-exporter . ; \
    fi

FROM alpine:3.19

COPY --from=builder /go/bin/immich-exporter /go/bin/immich-exporter

WORKDIR /go/bin

CMD ["/go/bin/immich-exporter"]
