FROM golang

ARG APPNAME

WORKDIR /

ADD $APPNAME /go_app

EXPOSE 30000

RUN go build -v -o ./jenkins-app

CMD ["./jenkins-app"]