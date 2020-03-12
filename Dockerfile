FROM golang:1.9
LABEL maintainer "Wls <wanglishuai_210@sina.com>"

RUN mkdir /data
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

CMD ["go-wrapper", "run"]

COPY . /go/src/app
RUN go-wrapper download
RUN go-wrapper install

VOLUME ["/data"]
ENTRYPOINT ["/go/bin/app", "-dbDir=/data"]

EXPOSE 8080

