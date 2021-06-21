FROM golang:1.16

WORKDIR /app

RUN go mod init example.com/hello
COPY ./code/ ./code/

RUN go mod tidy