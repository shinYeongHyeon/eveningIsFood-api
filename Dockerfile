#Dockerfile
FROM golang:1.16-alpine

WORKDIR /app
COPY docker .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex

# 9: HotReload

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]