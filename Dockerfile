FROM golang:1.20-alpine

# タイムゾーン設定のためのパッケージをインストール
RUN apk add --no-cache tzdata curl git bash

# airのインストール
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

WORKDIR /go/src/app

# タイムゾーンを設定
ENV TZ=Asia/Tokyo

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# wait-for-it.shスクリプトを追加
COPY scripts/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

CMD ["air"]
