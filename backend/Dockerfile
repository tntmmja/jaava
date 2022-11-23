



# see oli forumis
FROM golang:1.18

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]


# see on ka mingi variant
# FROM golang:1.18-alpine AS builder
# RUN apk update && apk add git && apk add build-base
# RUN mkdir /build
# ADD ./ /build
# WORKDIR /build

# RUN go get github.com/gorilla/websocket
# RUN go get github.com/mattn/go-sqlite3
# RUN go get github.com/satori/go.uuid
# RUN go get golang.org/x/crypto

# RUN go build -o real-time-forum

# CMD ["./real-time-forum"]
