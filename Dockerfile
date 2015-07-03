FROM google/golang
MAINTAINER Sakeven "sakeven.jiang@daocloud.io"

# Build app
WORKDIR /gopath/app
ENV GOPATH /gopath/app
ADD . /gopath/app/src/golang-redis-sample

RUN go get -t golang-redis-sample
RUN go install golang-redis-sample

EXPOSE 80
CMD ["/gopath/app/bin/golang-redis-sample"]
