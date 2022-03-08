FROM golang:1.17-alpine

WORKDIR /app    

COPY . .

RUN go get 

EXPOSE 8080

CMD make run