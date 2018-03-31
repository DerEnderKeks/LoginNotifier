FROM alanfranz/fwd-centos-7:latest
LABEL maintainer="DerEnderKeks"
RUN yum install -y wget git

ARG GOLANG_VERSION

RUN set -eux; \
	wget -O go.tgz "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz"; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz;

RUN wget -O upx.tar.xz "https://github.com/upx/upx/releases/download/v3.94/upx-3.94-amd64_linux.tar.xz"; \
    tar --wildcards --strip-components=1 -C /usr/local/bin -xf upx.tar.xz */upx; \
    rm upx.tar.xz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"