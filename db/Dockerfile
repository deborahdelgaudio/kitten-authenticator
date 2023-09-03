FROM golang:1.20

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY ./migrations/* ./migrations/
COPY Makefile .
COPY run.sh .

RUN chmod +x run.sh

CMD [ "sh", "run.sh" ]