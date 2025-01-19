FROM golang:1.23 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]