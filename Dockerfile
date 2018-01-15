FROM node:8.8.1

RUN mkdir -p /app
WORKDIR /app
COPY . /app
