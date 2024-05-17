FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN go mod download

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ./server main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]