FROM alpine:latest
WORKDIR /src/app
COPY ./ports.json .

ARG file= /src/app/ports.json

ADD main /
EXPOSE 8080
CMD ["./main","run"]