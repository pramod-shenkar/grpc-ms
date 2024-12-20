FROM golang:latest as builder
ARG servicename

WORKDIR /usr/src/app
COPY . .
# RUN go mod tidy

RUN ls
RUN go build -o main ./services/order/cmd/main.go

# FROM scratch
# COPY --from=builder /usr/src/app/main ./main
CMD ["./main"]

# CMD ["sleep", "5000000"]