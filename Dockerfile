FROM golang:alpine

WORKDIR /bin
COPY . /bin

RUN go build main.go

CMD [ "./main"]

EXPOSE 8080
