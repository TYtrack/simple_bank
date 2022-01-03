# Build Stage
FROM golang:1.17.5-alpine3.15 AS builder

WORKDIR /app
COPY . .

RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz
        

# Run Stage
FROM alpine:3.15 
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration


EXPOSE 9999

# CMD 和ENTRYPOINT一起使用的时候，CMD只是后者的参数
CMD ["/app/main"] 
ENTRYPOINT [ "/app/start.sh" ]