FROM golang:alpine

WORKDIR /app/days-remaining

RUN apk add --no-cache git tzdata 

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o ./out/days-remaining .

EXPOSE 8080

CMD ["./out/days-remaining"]