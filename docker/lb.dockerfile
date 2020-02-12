FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/JesseyBland/project-1/
COPY . .
EXPOSE 6061
RUN go get gopkg.in/yaml.v2
RUN go build ./cmd/loadbalancer

ENTRYPOINT [ "./loadbalancer", "--port", "6061" ]