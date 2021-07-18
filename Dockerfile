#Dockerfile
FROM golang:1.16-alpine

WORKDIR /app
COPY docker .

RUN apk update && \
    apk add git && \
    go get github.com/cespare/reflex &&
    go get github.com/gorilla/mux &&
    go get github.com/labstack/gommon &&
    go get -u gorm.io/gorm &&
    go get -u gorm.io/driver/postgres

# 9: HotReload

EXPOSE 9999
CMD ["reflex", "-c", "reflex.conf"]