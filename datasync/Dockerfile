FROM golang:1.15-alpine3.12

RUN mkdir /app
ADD . /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

## Add this go mod download command to pull in any dependencies
RUN go mod download

COPY *.go ./

RUN go build -o datasync
RUN chmod +x /app/datasync

# EXPOSE 8090

CMD [ "/app/datasync" ]
