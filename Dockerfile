FROM golang:1.19.2-alpine AS builder

RUN mkdir build
WORKDIR /build

COPY go.mod go.sum ./
COPY *.go ./
COPY scripts/out ./scripts/out
RUN go mod download
RUN apk update && apk add build-base
RUN go build main.go handlers.go vocab.go 

FROM alpine
COPY --from=builder /build/main /
COPY --from=builder /build/scripts/out /scripts/out
RUN apk update && apk add ffmpeg

CMD ["/main"]