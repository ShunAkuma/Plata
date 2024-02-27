FROM golang:latest

WORKDIR /app


COPY redis.conf /shared/redis.conf
COPY redis.log /shared/redis.conf
# Copies everything from your root directory into /app
COPY . .
 
# Installs Go dependencies
RUN go mod download
 
# Builds your app with optional configuration
RUN go build -o /backend ./cmd/main.go
 
# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
CMD ["/backend"]


# docker run --name backend -p 8080:8080 17f92b874e5fD