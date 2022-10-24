
FROM golang:1.19.2-alpine AS builder

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o ./vvgo

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/vvgo .

CMD ["./vvgo"]