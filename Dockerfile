FROM golang:1.19.2-alpine3.16

LABEL org.opencontainers.image.source="https://github.com/PascalRoose/elite1337bot"

ENV TGBOT_TOKEN "insert-token-here"

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]