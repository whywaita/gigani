FROM golang:1.16 AS builder

WORKDIR /go/src/github.com/whywaita/gigani

COPY . .
RUN go build .

FROM alpine

RUN apk update \
  && apk add --no-cache ca-certificates \
  && update-ca-certificates 2>/dev/null || true

COPY --from=builder /go/src/github.com/whywaita/gigani/gigani /app

CMD ["/app"]
