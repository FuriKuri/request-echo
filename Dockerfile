FROM golang:1.12
ADD . /go/src/github.com/furikuri/request-echo
WORKDIR /go/src/github.com/furikuri/request-echo
RUN go get ./
RUN go build

FROM alpine:3.10
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
WORKDIR /root/
COPY --from=0 /go/bin/request-echo .
EXPOSE 80
ENTRYPOINT ["/root/request-echo"]
