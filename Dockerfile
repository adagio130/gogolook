FROM golang:1.22-alpine AS builder
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY . .

RUN go env -w CGO_ENABLED=1
RUN go mod tidy

RUN go test -v ./...
RUN go build -o main .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .

CMD ["./main"]
