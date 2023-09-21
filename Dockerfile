FROM golang:1.19
ENV GO111MODULE=on
ENV CGO_ENABLED 0
COPY . /service
WORKDIR /service
RUN go build -o rest-api 

CMD ./rest-api
