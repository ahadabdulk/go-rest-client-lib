
FROM golang:latest

LABEL maintainer="Abdul Ahad <ahad.abdulk@gmail.com>"

WORKDIR /app

RUN go get -u github.com/stretchr/testify/assert

COPY . .

ENTRYPOINT ["go","test","-v", "-tags=integration"]