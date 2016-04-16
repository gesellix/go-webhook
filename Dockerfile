FROM alpine:edge
MAINTAINER Tobias Gesellchen <tobias@gesellix.de> (@gesellix)
EXPOSE 3003

ENV GOPATH /go
ENV APPPATH $GOPATH/src/github.com/gesellix/go-webhook
COPY . $APPPATH

RUN apk add --update -t build-deps go git mercurial libc-dev gcc libgcc \
    && cd $APPPATH && go get -d && go build -o /bin/go-webhook \
    && apk del --purge build-deps && rm -rf $GOPATH
    && chmod 755 /bin/go-webhook

ENTRYPOINT [ "/bin/go-webhook" ]
CMD [ "" ]
