#Dockerfile
FROM golang:1.18-alpine

WORKDIR /app

RUN apk update && \
    apk add git && \
    go install github.com/cespare/reflex@latest

# 8 Line: Hot-reload

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]
