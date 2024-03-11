FROM golang:1.22.1

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o ./main ./cmd/go-pokemon-challenge/main.go

ENTRYPOINT ["./main"]