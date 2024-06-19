FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum .env ./
RUN go mod download

COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN CGO_ENABLED=0 GOOS=linux swag init && go build -o lms

FROM alpine:3.20

COPY --from=builder /app/docs /app/lms /app/.env /

ENTRYPOINT ["/lms"]