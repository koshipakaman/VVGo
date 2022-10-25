FROM golang:1.19.2-alpine AS builder

RUN mkdir build
WORKDIR /build

COPY go.mod go.sum ./
COPY *.go ./
RUN go mod download
RUN apk update && apk add build-base
RUN go build -o vvgo 

FROM alpine
COPY --from=builder /build/vvgo /

CMD ["/vvgo"]