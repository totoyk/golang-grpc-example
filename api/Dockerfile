FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@v1.61.1

COPY . ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]

EXPOSE 50051
