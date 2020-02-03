FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/JesseyBland/project-1
COPY . .
EXPOSE 8888
RUN go build ./cmd/Tcp/Server

ENTRYPOINT [ "./Server", "--port", "8888" ]