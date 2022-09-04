# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

ARG APP_ENV
ENV APP_ENV $APP_ENV

WORKDIR /app

COPY . .
RUN go mod download

RUN cd cmd/TravelMaster && go build -o /app/travel_master && cd /app

EXPOSE 8080

CMD [ "/app/travel_master" ]