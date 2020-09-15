FROM golang:1.14-stretch as builder_base

WORKDIR /ydata-crd

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux

COPY go.mod .
COPY go.sum .

RUN go mod download

# STAGE 1: Build binaries
FROM builder_base as builder
COPY . /ydata-crd
WORKDIR /ydata-crd

RUN go build -a -installsuffix cgo -o server github.com/victoriavilasb/ydata-crd/cmd/server
ENTRYPOINT /go/bin/server