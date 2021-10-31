FROM golang:alpine

WORKDIR /go/src/gin-myboot
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN go generate && go env && go build -o server .

FROM alpine:latest
LABEL MAINTAINER="westhack@limaopu.com"

WORKDIR /go/src/gin-myboot

COPY --from=0 /go/src/gin-myboot ./

EXPOSE 8888

ENTRYPOINT ./server -c config.docker.yaml
