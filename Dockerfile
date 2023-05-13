FROM golang:1.20-alpine

WORKDIR /app
COPY . /app

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air && \
    go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
