FROM scratch

COPY ./server_bin /server
COPY ./insecure /insecure

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["/server"]
