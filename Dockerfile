FROM golang:1.20-alpine

WORKDIR /app
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
