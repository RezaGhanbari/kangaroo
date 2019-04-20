# consignment-service/Dockerfile
FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /app/kangaroo-service-consignment

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shippy-service-consignment

FROM alpine:latest

# Security related package, good to have.
RUN apk --no-cache add ca-certificates
RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/kangaroo-service-consignment/kangaroo-service-consignment .

CMD ["./kangaroo-service-consignment"]