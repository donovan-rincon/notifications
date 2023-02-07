FROM golang:1.16-alpine
WORKDIR /notifications


COPY . .

RUN go mod download


WORKDIR /notifications/cmd

RUN go build main.go

EXPOSE 80

CMD [ "./main" ]