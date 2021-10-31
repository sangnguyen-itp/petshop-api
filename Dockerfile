FROM golang:alpine

RUN yum update -y
RUN mkdir -p bin

COPY . /bin
WORKDIR /bin

RUN go build -o petshop-api

CMD [ "/petshop-api"]
EXPOSE 8080
