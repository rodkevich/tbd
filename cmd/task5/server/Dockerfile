# build stage
FROM golang:1.16 AS builder
LABEL maintainer="ivan.rodkevich"

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

#RUN go get -d -v github.com/go-playground/validator
#RUN go get -d -v github.com/google/uuid
#RUN go get -d -v github.com/gorilla/mux
#RUN dpkg-reconfigure ca-certificates

WORKDIR /go/src/github.com/rodkevich/tbd
COPY . .
RUN go mod init
RUN go get ./...
WORKDIR /go/src/github.com/rodkevich/tbd/cmd/task5/server

RUN go build -ldflags="-s -w" ./server.go

# server stage
FROM scratch
LABEL maintainer="ivan.rodkevich/app"

COPY --from=builder /go/src/github.com/rodkevich/tbd/cmd/task5/server .
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 12300

CMD ["./server"]
