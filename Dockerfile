FROM golang:1.23

WORKDIR /app

COPY . .

EXPOSE 3000

CMD ["go", "run", "cmd/main.go"]
