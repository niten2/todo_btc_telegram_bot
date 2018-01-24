FROM ubuntu:14.04

RUN mkdir -p /app-telegram
WORKDIR /app-telegram
COPY . /app-telegram
