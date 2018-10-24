FROM golang:1.11

ADD ./ /go/src/centre
WORKDIR /go/src/centre

# Install dependecies
RUN go get -u google.golang.org/grpc
RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get github.com/lib/pq

RUN export PATH=$PATH:$GOPATH/bin


CMD go run *.go