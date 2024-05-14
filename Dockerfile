FROM golang:1.22 AS build

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o shifu-helloworld .

CMD ["./shifu-helloworld"]
