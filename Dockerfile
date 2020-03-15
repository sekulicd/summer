FROM golang:1.13.0-stretch AS builder

ENV GO111MODULE=on

RUN make build

WORKDIR /dist

RUN cp summer ./dist/summer



