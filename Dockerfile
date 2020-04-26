FROM golang:alpine as builder

ENV WD=/go/src/github.com/monitoring-test/sample-service

COPY . $WD
WORKDIR $WD
RUN go build -o application .

FROM alpine

ENV WD=/go/src/github.com/monitoring-test/sample-service
COPY --from=builder $WD/application /app/application
WORKDIR /app
ENTRYPOINT ["./application"]