FROM golang:1.11.2-alpine3.8 as builder

RUN mkdir $GOPATH/src/cyphernode_lunanode_proxy

WORKDIR $GOPATH/src/cyphernode_lunanode_proxy

COPY main.go .
RUN go build -o /tmp/cyphernode_lunanode_proxy -ldflags="-s -w"

FROM alpine:3.8

RUN apk add --no-cache ca-certificates

COPY --from=builder /tmp/cyphernode_lunanode_proxy /usr/bin

WORKDIR /

COPY tsconf.json /tsconf.json

EXPOSE 9999

CMD ["/usr/bin/cyphernode_lunanode_proxy"]
