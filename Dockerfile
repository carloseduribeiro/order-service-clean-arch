FROM golang:1.22.0 AS builder

WORKDIR /app
COPY . .

RUN go get -v
RUN GOOS=linux CGO_ENABLED=0 go build -C /app/cmd/orderSystem -ldflags="-w -s" -o runner main.go wire_gen.go
RUN chmod +x /app/cmd/orderSystem/runner

FROM scratch

WORKDIR /app
COPY --from=builder /app/cmd/orderSystem/runner .
COPY ./cmd/orderSystem/.env .env

CMD ["/app/runner"]