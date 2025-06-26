FROM golang:1.22

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy
RUN go build -o app .

CMD ["./app"]

