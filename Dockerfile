FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/polygon-edge

FROM alpine:3.18 as runner
COPY --from=builder /app/polygon-edge /app/polygon-edge
COPY --from=builder /app/genesis.json /app/genesis.json
EXPOSE 8545 9632 1478 5001 30301 30302
ENTRYPOINT ["/app/polygon-edge"]