FROM golang:1.19-alpine3.17 AS builder

WORKDIR /app

COPY . .

RUN go get -d -v ./src/ && \
    go build -o /go/bin/immich-exporter ./src 

FROM alpine:3.17

COPY --from=builder /go/bin/immich-exporter /go/bin/immich-exporter

CMD ["/go/bin/immich-exporter"]
