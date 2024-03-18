FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# install goose
RUN go get -u github.com/pressly/goose/cmd/goose

# build go app
RUN go mod download
RUN go build -o vk-vilmoteca ./cmd/main.go

CMD ["./vk-vilmoteca"]