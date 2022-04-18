FROM golang:1.18-alpine3.15

WORKDIR /opt/app

COPY . .

RUN go build -o cli-proposal /opt/app/cmd/cli/main.go

CMD sleep 1d