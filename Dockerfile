FROM golang:1.15.0-alpine3.12 as builder
RUN adduser -u 1000 -D app-runner
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /opt/app
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o main .

FROM alpine as final

RUN apk add --no-cache tzdata
ENV TZ="Asia/Shanghai"
WORKDIR /opt/app
COPY --from=builder /opt/app/main .
EXPOSE 8090
USER app-runner
CMD ["/opt/app/main"]