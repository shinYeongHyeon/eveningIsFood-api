#Dockerfile
FROM golang:1.17-alpine

WORKDIR /app

# 9 Line: Hot-reload
RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]