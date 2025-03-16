FROM golang:1.24.1-alpine AS builder
ENV GO111MODULE=on
RUN apk add ca-certificates
WORKDIR /app
ADD . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -extldflags "-static"'  -o netdemo ./cmd

FROM scratch
EXPOSE 80 
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/netdemo /bin/netdemo
ENTRYPOINT ["/bin/netdemo"]