FROM golang:1.19-alpine as builder
LABEL maintainer="argos team"
WORKDIR /app
COPY go.mod go.sum ./
RUN apk update && apk add gcc openssl-libs-static zlib-static zstd-libs libsasl lz4-dev lz4-static zstd-static libc-dev musl-dev
RUN go mod download
COPY . .
RUN GOGOARCH=amd64 GOOS=linux go build -tags musl -ldflags '-w -extldflags "-static"' -a -installsuffix cgo -o main cmd/api/main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]