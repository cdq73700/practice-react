FROM golang:1.19-alpine

WORKDIR /backend

COPY go.mod .
COPY go.sum .

RUN apk update &&  apk add git
RUN go install github.com/cosmtrek/air@latest
 
# go modからパッケージをダウンロード
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]