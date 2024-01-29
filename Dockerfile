FROM golang:alpine3.18 as builder

WORKDIR /opt

ENV LISTEN_PORT=$LISTEN_PORT

COPY . .

RUN go build -o counter-service main.go

FROM alpine:latest

COPY --from=builder /opt/counter-service .

CMD ./counter-service