# - builder
FROM golang:1.11-alpine as builder
RUN apk update && apk upgrade && \
    apk --no-cache --update add bash git make gcc g++ libc-dev
WORKDIR /go/src/github.com/ftier-stock
ENV GO111MODULE=on
COPY . .
RUN go mod vendor -v && go build -o engine app/main.go

# - distribution
FROM alpine:latest
RUN apk update && apk upgrade && \
    apk --no-cache --update add ca-certificates tzdata && \
    mkdir /ftier-stock && mkdir /app
WORKDIR /ftier-stock

EXPOSE 3001

COPY --from=builder /go/src/github.com/ftier-stock/engine /app
COPY --from=builder /go/src/github.com/ftier-stock/app.config* ./
RUN ls -lh
CMD /app/engine