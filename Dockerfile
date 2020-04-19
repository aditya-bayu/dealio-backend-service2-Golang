FROM golang:1.14.1

WORKDIR /app
COPY code/ /app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/jinzhu/gorm

RUN go build -o application

CMD [ "./application" ]