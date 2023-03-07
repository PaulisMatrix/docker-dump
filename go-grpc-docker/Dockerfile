FROM ubuntu:20.04

# install required base packages
RUN apt-get update && apt-get -y upgrade && apt-get install -y git make curl autoconf automake libtool g++ unzip

# install go
RUN curl -O https://storage.googleapis.com/golang/go1.18.linux-amd64.tar.gz

# extract
RUN tar -C /usr/local -xzf go1.18.linux-amd64.tar.gz

# set stuff for go
RUN mkdir -p /root/go/bin /root/go/src /usr/local/go
ENV LC_CTYPE='C.UTF-8' GOPATH="/root/go" PATH="${PATH}:/usr/local/go/bin:/usr/local/linux-amd64:/root/bin:/root/go/bin"

# get the protobuf and protoc
RUN curl -L -H 'Accept: Application/octet-stream' -o protobuf.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.8.0/protoc-3.8.0-linux-x86_64.zip
RUN unzip -qq protobuf.zip -d protoc3
RUN mv protoc3/bin/* /usr/local/bin/
RUN mv protoc3/include/* /usr/local/include/

# go plugins for grpc
RUN  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN  export PATH="$PATH:$(go env GOPATH)/bin"

# for app
WORKDIR /usr/src/app

COPY . /usr/src/app/

# switch to proto dir
WORKDIR /usr/src/app/my_albums

# generate protos
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./*.proto

RUN go mod download && go mod tidy

# switch to main dir
WORKDIR /usr/src/app 

RUN CGO_ENABLED=0 go build -o go-grpc-docker main.go

CMD ["/go-grpc-docker"]