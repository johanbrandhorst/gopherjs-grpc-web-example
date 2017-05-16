FROM ubuntu:14.04

COPY ./nginx /nginx
COPY ./nginx.conf /nginx.conf

EXPOSE 8081

RUN ["mkdir", "-p", "/logs"]

ENTRYPOINT ["/nginx", "-p", "/", "-c", "/nginx.conf"]
