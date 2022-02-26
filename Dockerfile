#Dockerfile
FROM golang:1.17-alpine

WORKDIR /app

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex && \
    go get github.com/gofiber/fiber/v2

# 8 Line: Hot-reload
# 9 Line: Web Framework "fiber"

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]