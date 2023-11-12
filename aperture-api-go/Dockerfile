FROM golang:1.19

WORKDIR /go/src/app

COPY . .

EXPOSE 8080

RUN go build -o main main.go

CMD ["./main"]

