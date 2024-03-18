FROM golang:alpine3.19 AS builder

WORKDIR /build

COPY go.mod .

COPY . .

RUN go build -o laAPI ./cmd/service/main.go

FROM ubuntu:20.04

ENV APP_DIR /build

WORKDIR $APP_DIR

RUN apt-get update \
    && groupadd -r web \
    && useradd -d $APP_DIR -r -g web web \
    && chown web:web -R $APP_DIR \
    && apt-get install -y netcat-traditional \
    && apt-get install -y acl

COPY --from=builder /build/laAPI $APP_DIR/laAPI
COPY --from=builder /build/deploy/scripts/local-laAPI-start.sh $APP_DIR/local-laAPI-start.sh
COPY --from=builder /build/deploy/env/.env.local $APP_DIR//deploy/env/.env.local

RUN setfacl -R -m u:web:rwx $APP_DIR/local-laAPI-start.sh

USER web

ENTRYPOINT ["bash", "local-laAPI-start.sh"]