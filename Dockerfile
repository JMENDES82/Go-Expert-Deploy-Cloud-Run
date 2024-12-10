FROM golang:1.21.4 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=builder /app/server /server
EXPOSE 8080
CMD ["/server"]
