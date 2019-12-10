FROM golang:latest

WORKDIR $GOPATH/src/github.com/material-seminario-golang/deployment
ADD . $GOPATH/src/github.com/material-seminario-golang/deployment

RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/swaggo/swag/cmd/swag

RUN govendor init
RUN govendor sync
RUN swag init

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
