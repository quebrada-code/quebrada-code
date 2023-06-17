#FROM golang:alpine AS builder
#WORKDIR /build
#ADD go.mod .
#RUN go install ./...
#COPY . .

#RUN go build -o solution-worker ./cmd/main.go
#
#FROM alpine
#WORKDIR /build
#COPY --from=builder /build/solution-worker /build/solution-worker
#CMD ["./solution-worker"]

FROM golang:alpine AS builder

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev librdkafka-dev pkgconf

WORKDIR /build

# Fetch dependencies
COPY go.mod go.sum ./
# Build
COPY . ./
RUN CGO_ENABLED=1 go build -tags muls  ./...
# Create final image
FROM alpine
WORKDIR /
COPY --from=builder /build/solution-worker .
CMD ["./solution-worker"]