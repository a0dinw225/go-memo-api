package routes

import (
	"go-memo-api/internal/config"
	"go-memo-api/internal/controllers"
	"go-memo-api/internal/middlewares"
	"go-memo-api/internal/repositories"
	"go-memo-api/internal/services"
	"go-memo-api/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, cfg config.Config) {
	db := utils.SetupDatabaseConnection(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	tagRepo := repositories.NewTagRepository(db)
	tagService := services.NewTagService(tagRepo)
	tagController := controllers.NewTagController(tagService)

	v1 := r.Group("/api/v1")

	// 認証ミドルウェア
	v1.Use(middlewares.AuthMiddleware(cfg.AppKey, cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))

	log.Println("Routes have been successfully set up")

	v1.DELETE("/tag/delete/:id", tagController.DeleteTag)
}
