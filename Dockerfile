FROM golang:1.19.2-alpine3.16

ENV TGBOT_TOKEN "insert-token-here"

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]