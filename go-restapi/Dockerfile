FROM golang:1.21

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build -o api .

EXPOSE 3000

CMD ["./api"]