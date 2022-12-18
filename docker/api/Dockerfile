FROM golang:1.19-alpine

ENV WORKDIR /go/src

WORKDIR $WORKDIR


COPY ./src .

RUN go get -u github.com/cosmtrek/air
RUN go build -o /go/bin/air github.com/cosmtrek/air
CMD ["air", "-c", ".air.toml"]