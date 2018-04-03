FROM golang

EXPOSE 8888

RUN mkdir /go/src/user-api-endpoint

ADD . /go/src/user-api-endpoint

WORKDIR /go/src/user-api-endpoint

RUN go get -d -v ./...

RUN go build -o main.go

RUN chmod +x user-api-endpoint

ENTRYPOINT ["./user-api-endpoint"]
