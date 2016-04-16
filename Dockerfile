FROM alpine:edge
MAINTAINER Tobias Gesellchen <tobias@gesellix.de> (@gesellix)
EXPOSE 3003

ENV GOPATH /go
ENV APPPATH $GOPATH/src/github.com/gesellix/go-webhook
COPY . $APPPATH

RUN apk add --update -t build-deps go git mercurial libc-dev gcc libgcc \
    && cd $APPPATH && go get -d && go build -o /bin/go-webhook cli/cli.go \
    && apk del --purge build-deps && rm -rf $GOPATH

ENTRYPOINT [ "/bin/go-webhook" ]
CMD [ "-listen-address=0.0.0.0:3003" ]
