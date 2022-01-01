# Build Stage
FROM golang:1.17.5-alpine3.15 AS builder

WORKDIR /app
COPY . .

RUN go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
RUN go build -o main main.go

# Run Stage
FROM alpine:3.15 
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 9999
CMD ["/app/main"] 