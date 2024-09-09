FROM golang:1.23-alpine3.20 AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN GOOS=linux go build -o example-go-service ./cmd/api/main.go



FROM alpine:3.20 AS final

COPY --from=builder /app/example-go-service /app/example-go-service
COPY ./config.yaml /app/config.yaml

WORKDIR /app
CMD [ "./example-go-service" ]
