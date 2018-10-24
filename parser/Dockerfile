FROM golang:1.11

ADD ./ /go/src/parser
WORKDIR /go/src/parser

# Install dependecies
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/mholt/archiver/cmd/archiver

RUN export PATH=$PATH:$GOPATH/bin

CMD go run *.go