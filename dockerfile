FROM golang:alpine as builder
ARG servicename

WORKDIR /usr/src/app
COPY . .
# RUN go mod tidy

RUN go build -o main ./services/${servicename}/cmd/main.go

# FROM sratch
# COPY --from=builder /usr/src/app/order ./order
CMD ["./main"]

# CMD ["sleep", "5000000"]