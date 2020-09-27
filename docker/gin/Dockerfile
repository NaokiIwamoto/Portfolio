FROM golang:latest

WORKDIR /go/src/github.com/api

RUN apt-get update \
  && apt-get install -y \
  git \
  zip \
  unzip \
  vim

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

CMD ["air"]