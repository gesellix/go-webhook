FROM alpine:3.11 AS builder
LABEL builder=true

ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV APPPATH /app

RUN adduser -DH user
RUN apk add --update -t build-deps go git mercurial libc-dev gcc libgcc
COPY . $APPPATH
RUN cd $APPPATH && go get -d \
 && go test -short ./... \
 && go build \
    -a \
    -ldflags '-s -w -extldflags "-static"' \
    -o /bin/go-webhook \
    cli/cli.go

FROM scratch
LABEL maintainer="Tobias Gesellchen <tobias@gesellix.de> (@gesellix)"

EXPOSE 3003

ENTRYPOINT [ "/bin/go-webhook" ]
CMD [ "-listen-address=0.0.0.0:3003" ]

COPY --from=builder /etc/passwd /etc/passwd
USER user

COPY --from=builder /bin/go-webhook /bin/go-webhook
