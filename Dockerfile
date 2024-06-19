FROM golang:1.20

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN go build -o main .

CMD ["./main"]
