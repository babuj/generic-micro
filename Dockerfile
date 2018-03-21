FROM alpine
LABEL author="babuj"

COPY generic-micro /app/generic-micro
ENTRYPOINT [ "/app/generic-micro" ]

# FROM golang:alpine
# LABEL author=babuj

# ENV SOURCES /go/src/github.com/babuj/generic-micro
# COPY . ${SOURCES}
# RUN cd ${SOURCES} && CGO_ENABLED=0 go install

# ENV PORT 8080
# ENV NAME Echo
# EXPOSE 8080
# ENTRYPOINT [ "/go/bin/generic-micro" ]