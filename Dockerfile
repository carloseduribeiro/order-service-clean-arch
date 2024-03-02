FROM golang:1.22.0 AS builder

WORKDIR /app
COPY . .

RUN go get -v
RUN CGO_ENABLED=0 go build -C /app/cmd/orderSystem -o runner main.go wire_gen.go
RUN chmod +x /app/cmd/orderSystem/runner

FROM scratch

WORKDIR /app
COPY --from=builder /app/cmd/orderSystem/runner ./runner
COPY ./cmd/orderSystem/.env .env

ENTRYPOINT ["/app/runner"]