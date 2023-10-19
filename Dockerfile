FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go build -o /go/bin/immich-exporter ./src 


FROM alpine:3.18

COPY --from=builder /go/bin/immich-exporter /go/bin/immich-exporter
COPY package.json /go/bin/

WORKDIR /go/bin

CMD ["/go/bin/immich-exporter"]