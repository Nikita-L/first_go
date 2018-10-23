FROM golang:1.11

ADD ./ /go/src/parser
WORKDIR /go/src/parser

# Install dependecies
RUN go get -u google.golang.org/grpc


CMD go run ./main.go