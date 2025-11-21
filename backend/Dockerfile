# Go Image
FROM golang:1.23 AS builder

WORKDIR /app

RUN go mod init github.com/kalebhilde/SortingAlgoBenchmarkApp || true

# COPY go.mod go.sum
# RUN go mod download

COPY . .

RUN go mod init backend || true
RUN go mod tidy

# Build binary
RUN go build -o server main.go

# Final stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]