FROM golang:alpine

WORKDIR /app


COPY redis.conf /shared/redis.conf
COPY . .
 
RUN go mod download
 
RUN go build -o /backend ./cmd/main.go
 
EXPOSE 8080
 
CMD ["/backend"]