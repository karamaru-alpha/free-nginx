FROM golang:1.18-alpine3.15 as builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
WORKDIR /go/src/github.com/karamaru-alpha/free-nginx

COPY . ./
RUN go build -o app

FROM alpine:3.15
WORKDIR /root/
COPY --from=builder /go/src/github.com/karamaru-alpha/free-nginx .

ENTRYPOINT ["./app"]
