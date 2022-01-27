FROM golang:1.15-alpine3.12

RUN mkdir /app
ADD . /app

WORKDIR /app

## Add this go mod download command to pull in any dependencies
RUN go mod download

RUN go build -o datasync-service

CMD [ "/datasync-service" ]
