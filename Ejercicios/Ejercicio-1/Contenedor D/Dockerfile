FROM golang:latest

RUN mkdir app_code

WORKDIR /app_code

COPY . .

RUN go get go.mongodb.org/mongo-driver/mongo

RUN go get -u github.com/go-chi/chi

