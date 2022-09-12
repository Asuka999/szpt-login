FROM debian:stable-slim


WORKDIR /

COPY ./build/szpt_login_linux_amd64_linux .


EXPOSE 50051

ENTRYPOINT ./szpt_login_linux_amd64_linux









