FROM alpine:3

LABEL maintainer="Mioto <yaku.mioto@gmail.com>"


RUN apk update && apk add --no-cache ca-certificates

COPY hlf-tools /usr/local/bin/hlf-tools

ENTRYPOINT ["hlf-tools"]