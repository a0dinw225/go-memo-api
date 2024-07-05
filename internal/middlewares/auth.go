package middlewares

import (
	"go-memo-api/internal/models"
	"go-memo-api/pkg/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ログイン認証ミドルウェア
// クッキーまたはAuthorizationヘッダーからトークンを取得し、データベースに存在するか確認する
// トークンが存在しないまたは無効である場合は401エラーを返す
func AuthMiddleware(secretKey string, dbUser, dbPassword, dbHost, dbPort, dbName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenString string
		var err error

		// クッキーからトークンを取得
		tokenString, err = c.Cookie("laravel_memo_app_sanctum_token")
		if err != nil {
			// クッキーが見つからない場合はヘッダーから取得を試みる
			tokenString = c.GetHeader("Authorization")
			if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
				tokenString = tokenString[7:]
			} else {
				log.Println("Unauthorized: Token not found in cookies or Authorization header")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Token not found in cookies or Authorization header"})
				c.Abort()
				return
			}
		}

		// tokenStringからBearerを削除
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// データベース接続をセットアップ
		db := utils.SetupDatabaseConnection(dbUser, dbPassword, dbHost, dbPort, dbName)

		// トークンからユーザーIDを取得
		var tokenRecord models.PersonalAccessToken
		err = db.Where("token = ?", tokenString).First(&tokenRecord).Error
		if err != nil {
			log.Printf("Unauthorized: Token not found in database. Error: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized", "message": "Token not found in database"})
			c.Abort()
			return
		}

		// コンテキストにユーザーIDをセット
		userId := tokenRecord.TokenableID
		c.Set("userId", userId)

		c.Next()
	}
}
