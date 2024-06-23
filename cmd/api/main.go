package main

import (
	"flag"
	"fmt"
	"go-memo-api/internal/config"
	"go-memo-api/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// ポートフラグの追加
	port := flag.String("port", "8082", "ポート番号")
	flag.Parse()

	// ログファイルを開く
	logFile, err := os.OpenFile("tmp/local.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	// ログの出力先をファイルに設定
	log.SetOutput(logFile)

	// .envファイルを読み込む
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	r := gin.Default()

	// ルートのセットアップ
	routes.SetupRoutes(r, cfg)

	log.Println("Routes have been successfully set up")

	// サーバーを開始
	err = r.Run(fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
