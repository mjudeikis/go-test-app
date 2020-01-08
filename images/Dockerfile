FROM docker.io/library/fedora
COPY ./static /static
COPY http /usr/local/bin/http
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/http"]
