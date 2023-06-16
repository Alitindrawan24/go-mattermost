FROM golang:1.19-alpine

RUN apk update && apk add --no-cache build-base

WORKDIR /app
COPY . /app

EXPOSE 8080

CMD ["make", "run"]