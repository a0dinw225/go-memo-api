package controllers

import (
	"log"
	"net/http"
	"strconv"

	"go-memo-api/internal/services"

	"github.com/gin-gonic/gin"
)

type TagController struct {
	tagService services.TagService
}

func NewTagController(tagService services.TagService) *TagController {
	return &TagController{tagService}
}

func (ctrl *TagController) DeleteTag(c *gin.Context) {
	log.Println("DeleteTag controller called")
	idStr := c.Param("id")
	tagID, err := strconv.Atoi(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tag ID"})
		return
	}

	// タグが削除されていないか確認
	isNotDeleted, err := ctrl.tagService.CheckTagNotDeleted(tagID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check tag"})
		return
	}
	if !isNotDeleted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tag is already deleted"})
		return
	}

	err = ctrl.tagService.DeleteTag(tagID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "タグを削除しました"})
}
