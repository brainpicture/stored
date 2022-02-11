FROM ubuntu:20.04

ADD tonlib-go tonlib-go
RUN apt update
ENV DEBIAN_FRONTEND=noninteractive
RUN apt -y install git gcc cmake
RUN apt -y install g++

RUN apt -y install wget
RUN wget -P /tmp https://go.dev/dl/go1.18beta2.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf /tmp/go1.18beta2.linux-amd64.tar.gz
RUN rm /tmp/go1.18beta2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

ENV LD_LIBRARY_PATH=/tonlib-go/v2/lib/linux
# ADD . /tonapi
ENV DEV 1
