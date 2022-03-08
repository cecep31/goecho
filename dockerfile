FROM golang:1.17.8-bullseye

WORKDIR /app    

COPY . .

COPY go get

EXPOSE 8080

CMD make run