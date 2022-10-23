
FROM golang:1.19.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o vvgo

CMD ["./vvgo"]