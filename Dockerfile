FROM golang:latest
LABEL authors="fanr"

COPY . /go/src/

WORKDIR "/go/src/"
RUN go env -w GO111MODULE=on \
  && go env -w GOPROXY=https://goproxy.cn,direct \
  && go env -w GOOS=linux \
  && go env -w GOARCH=amd64
RUN go mod tidy
RUN go build -o bibi