FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk add --no-cache python3 py3-pip
WORKDIR /root/
COPY --from=builder /app/main .
COPY public /public
EXPOSE 3000 8000
CMD python3 -m http.server 8000 --directory /public & ./main