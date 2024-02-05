FROM golang:1.21
LABEL maintainer="richardscull"

WORKDIR /ascii-art-web-dockerize

COPY . .
RUN go build -o main

EXPOSE 8080

CMD ["./main"]