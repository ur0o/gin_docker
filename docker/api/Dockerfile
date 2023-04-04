FROM golang:1.20

ENV WORKDIR /go/src

WORKDIR $WORKDIR


COPY ./src/go.mod .
COPY ./src/go.sum .

RUN go get -u github.com/cosmtrek/air
RUN go build -o /go/bin/air github.com/cosmtrek/air