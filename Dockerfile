FROM golang:1.20-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o schools ./cmd

RUN chmod 755 /app/schools

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/schools /app

CMD [ "/app/schools" ]