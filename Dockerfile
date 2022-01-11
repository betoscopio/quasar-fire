# builder image
FROM golang:1.17.6-alpine3.15 as builder
RUN mkdir -p /usr/local/go/src/quasar-fire 
COPY . /usr/local/go/src/quasar-fire/
WORKDIR /usr/local/go/src/quasar-fire
RUN go build -o quasar-fire-service

EXPOSE 8080

# generate clean, final image for end users
FROM alpine:3.15.0
COPY --from=builder /usr/local/go/src/quasar-fire/quasar-fire-service .
#TO change, modify the app or allow
EXPOSE 8080

# executable
ENTRYPOINT [ "/quasar-fire-service" ]