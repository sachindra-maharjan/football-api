FROM golang:1.15-alpine3.12

ENV GO111MODULE=on
ENV PORT=9000

RUN mkdir /app
# ADD . /app
COPY client /app

RUN ls

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

## Add this go mod download command to pull in any dependencies
RUN go mod download

RUN pwd
RUN ls

# COPY client ./
RUN cd client/httpclient

RUN go build -o httpclient

RUN pwd
RUN ls

## RUN chmod +x /app/datasync

# EXPOSE 8090

CMD [ "/app/httpclient" ]
