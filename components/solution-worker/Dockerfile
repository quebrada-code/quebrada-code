FROM golang:1.20-alpine3.17 AS builder


RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev librdkafka-dev pkgconf

WORKDIR "/code"
ADD . "/code"

RUN go build -tags musl -o solution-worker ./cmd/main.go

FROM alpine
WORKDIR /
COPY --from=builder /code/solution-worker .
CMD ["./solution-worker"]