FROM golang:1.25.0-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/cmd/checklist/exe /app/cmd/checklist

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/cmd/checklist/exe /app
ENV PROJECT_ROOT=/app
CMD ["/app/exe"]