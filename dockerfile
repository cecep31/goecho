FROM golang:1.16

WORKDIR /app    

COPY . .

COPY go get

EXPOSE 8080

CMD make run