FROM golang:1.22.1 AS build

WORKDIR /go/src

COPY . ./

RUN go mod download

RUN go build -o ./main ./cmd/go-pokemon-challenge/main.go

FROM  golang:1.22.1 AS start

COPY --from=build ./go/src/main ./
COPY --from=build ./go/src/.env ./

ENTRYPOINT ["./main"]