FROM golang:1.20-alpine

# 必要なパッケージをインストール
RUN apk add --no-cache tzdata curl git bash

WORKDIR /go/src/app

# タイムゾーンを設定
ENV TZ=Asia/Tokyo

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# wait-for-it.shスクリプトを追加
COPY scripts/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

CMD ["sh", "-c", "/wait-for-it.sh db:3306 && go run cmd/api/main.go"]
