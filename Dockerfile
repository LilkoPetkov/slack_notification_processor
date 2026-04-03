FROM golang:1.26

WORKDIR /app

COPY . .
RUN go mod tidy

RUN go build -o app main.go

EXPOSE 8081

CMD ["./app"]
