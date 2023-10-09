FROM golang:alpine3.18

WORKDIR /app

COPY . .

RUN go mod download

COPY *.go ./

RUN go build -o /Base64-api

EXPOSE 3000

CMD ["/Base64-api"]