FROM golang:1.15.3-alpine3.12

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build

CMD ["./tap-talk"]