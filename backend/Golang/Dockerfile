FROM golang:1.17.7-alpine

WORKDIR /backend

COPY go.mod .
COPY go.sum .

RUN apk update &&  apk add git
RUN go get github.com/cosmtrek/air@v1.29.0
 
# go modからパッケージをダウンロード
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]