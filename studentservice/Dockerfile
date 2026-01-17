FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app .

FROM golang:1.23-alpine
WORKDIR /app
COPY --from=builder app/app .
EXPOSE 5001
CMD ["./app"]