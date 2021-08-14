FROM golang
ADD . /go_app
WORKDIR /go_app
EXPOSE 30000
RUN go build -v -o /go_app/jenkins-app

CMD ["./jenkins-app"]
