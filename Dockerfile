FROM golang:latest
RUN mkdir -p /go/src/Gateway
ADD . /go/src/Gateway/
WORKDIR /go/src/Gateway/
RUN go get ./
RUN go build -o main .
CMD ["/go/src/Gateway/main"]