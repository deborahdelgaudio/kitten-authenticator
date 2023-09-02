FROM golang:1.20 as builder

WORKDIR /app

COPY src/ ./

RUN go mod download
RUN go build -v -o /kitten-server

EXPOSE 8080

CMD [ "/kitten-server" ]