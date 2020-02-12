FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/JesseyBland/project-1/
COPY . .
EXPOSE 6060
RUN go get gopkg.in/yaml.v2
RUN go build ./cmd/reverseproxy

ENTRYPOINT [ "./reverseproxy", "--port", "6060" ]





