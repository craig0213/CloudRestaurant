FROM golang:alpine

WORKDIR /Cloudrestaurant
ADD . /Cloudrestaurant
ENV GO111MODULE=on
RUN cd /Cloudrestaurant 
RUN go mod download 
RUN go build
EXPOSE 8080
ENTRYPOINT ./CloudRestaurant
