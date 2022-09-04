FROM golang:1.18


WORKDIR /szpt_login

COPY . .

ENV GOPROXY  "https://goproxy.cn"

CMD go mod tidy
CMD go build .

EXPOSE 50051

ENTRYPOINT ./szpt_login









