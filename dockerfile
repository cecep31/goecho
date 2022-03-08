FROM golang:1.17-alpine

WORKDIR /app    

COPY . .

RUN go get

RUN go build -o bin

EXPOSE 8080

CMD ./bin/goecho