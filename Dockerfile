FROM golang:1.18-alpine

ENV APP_ROOT=/usr/src/go-sample

RUN apk update && apk add make g++ sqlite-dev

RUN addgroup -S app && adduser -S -G app app
USER app

WORKDIR $APP_ROOT
COPY --chown=app:app . $APP_ROOT

RUN go build

EXPOSE 8080

CMD ["./go-sample"]
