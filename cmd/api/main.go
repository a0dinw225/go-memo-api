package main

import (
	"fmt"
	"go-memo-api/config"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.LoadConfig()

	// Ginを使った簡単なルートの定義
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Gormを使った簡単なデータベース接続
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	var db *gorm.DB
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database (attempt %d): %s\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database after 10 attempts: %s", err)
	}

	// テーブル名を取得してログに流す
	var tables []string
	db.Raw("SHOW TABLES").Scan(&tables)
	log.Println("Tables in the database:")
	for _, table := range tables {
		log.Println(table)
	}

	r.Run(":8080")
}
