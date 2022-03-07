FROM golang:1.14-buster

RUN go version
ENV GOPATH=/

COPY ./ ./


# install psql
RUN apt-get update

WORKDIR /var/golang


RUN build go app
RUN go mod download
RUN go build -o var/golang ./golang/golang-container.go


# CMD ["./var/golang]

CMD ping 127.0.0.1 -t

