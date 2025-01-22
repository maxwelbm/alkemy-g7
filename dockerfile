
FROM golang:1.23.3 AS builder


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


WORKDIR /app/cmd
RUN go build -o app .


FROM gcr.io/distroless/base


COPY --from=builder /app/cmd/app .


EXPOSE 8080


CMD ["./app"]