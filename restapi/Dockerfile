FROM golang:1.11
ENV APIPATH "/api/books"
ENV DBHOST "localhost:3306"
ENV DBPASS "1Vivek$ingh"

RUN mkdir /go/src/app
ADD . /go/src/app



WORKDIR /go/src/app

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/handlers

RUN go build -o main .

CMD ["/go/src/app/main"]
