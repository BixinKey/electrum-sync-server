#FROM golang:1.9
FROM golang:1.13
LABEL maintainer "Wls <wanglishuai_210@sina.com>"

RUN mkdir /data
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

#CMD ["go-wrapper", "run"]

COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
#RUN go-wrapper download
#RUN go-wrapper install

CMD ["app"]

VOLUME ["/data"]
ENTRYPOINT ["/go/bin/app", "-dbDir=/data"]

EXPOSE 8080

