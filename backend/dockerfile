FROM golang

COPY . /root/app

RUN go get github.com/gorilla/mux

CMD [ "go run main.go" ]