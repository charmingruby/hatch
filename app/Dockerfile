# Stage 1: Install dependencies
FROM golang:1.25 AS deps

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Stage 2: Build the application
FROM golang:1.25 AS builder

WORKDIR /app

COPY --from=deps /go/pkg /go/pkg
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/api/main.go

# Final stage: Run the application
FROM gcr.io/distroless/base-debian12:nonroot

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]
