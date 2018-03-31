FROM alanfranz/fwd-debian-jessie:latest
LABEL maintainer="DerEnderKeks"
RUN apt-get update; apt-get install -y wget git upx

ARG GOLANG_VERSION

RUN set -eux; \
	wget -q -O go.tgz "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz"; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz;

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"