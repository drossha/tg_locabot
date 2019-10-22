FROM golang:1.12-alpine as build

WORKDIR /go/src/tg_locbot
ADD . /go/src/tg_locbot

RUN apk add git

RUN CGO_ENABLED=0 GOOS=linux GO111MODULE=on go build -o /go/src/bin/tg_locbot

FROM alpine:latest

RUN apk add --no-cache bash ca-certificates

COPY --from=build /go/src/bin/tg_locbot .

ENV PORT=8080

CMD ["./tg_locbot"]

