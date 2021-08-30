# build stage
FROM golang:1.16 AS builder
LABEL maintainer="ivan.rodkevich"

ENV GO111MODULE=auto
ENV CGO_ENABLED=0
ENV GOOS=linux

#RUN go get -d -v github.com/google/uuid
#RUN dpkg-reconfigure ca-certificates

WORKDIR /go/src/github.com/rodkevich/tbd
COPY . .
WORKDIR /go/src/github.com/rodkevich/tbd/cmd/task5/app

RUN go build -ldflags="-s -w" ./app.go

# server stage
FROM scratch
LABEL maintainer="ivan.rodkevich/app"

COPY --from=builder /go/src/github.com/rodkevich/tbd/cmd/task5/app .
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 12300

CMD ["./app"]
