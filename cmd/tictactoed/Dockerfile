# Starting with the official golang/alpine image, tagged as our builder
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/JesseyBland/project-0
COPY . .
EXPOSE 8888
RUN go build ./cmd/tictactoed 

ENTRYPOINT [ "./tictactoed", "--port", "8888" ]