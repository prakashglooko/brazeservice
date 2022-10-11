# Build stage
FROM golang:1.19.2-alpine3.16 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/main .

EXPOSE 9090
CMD [ "/app/main"]