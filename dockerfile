FROM golang:1.18


WORKDIR /szpt_login

COPY . .

ENV GOPROXY  "https://goproxy.cn"

RUN go build -o ./szpt_login rpc/service/main.go

EXPOSE 50051

ENTRYPOINT ./szpt_login









