FROM golang:1.11-alpine

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ARG PROJECT_PACKAGE_NAME

ADD . $GOPATH/src/$PROJECT_PACKAGE_NAME

RUN apk add --no-cache ca-certificates \
        dpkg \
        gcc \
        git \
        musl-dev \
        curl \
    && mkdir -p "$GOPATH/src" "$GOPATH/bin" \
    && chmod -R 777 "$GOPATH" \
    && go get github.com/derekparker/delve/cmd/dlv

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR $GOPATH/src/$PROJECT_PACKAGE_NAME

RUN dep ensure
#RUN go build -o main .
