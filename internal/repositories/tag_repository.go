package repositories

import (
	"go-memo-api/internal/models"

	"gorm.io/gorm"
)

type TagRepository interface {
	GetTagByID(tagID int) (models.Tag, error)
	DeleteTag(tagID int) error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db}
}

func (r *tagRepository) GetTagByID(tagID int) (models.Tag, error) {
	var tag models.Tag
	result := r.db.First(&tag, tagID)
	return tag, result.Error
}

func (r *tagRepository) DeleteTag(tagID int) error {
	// ソフトデリートを使用して論理削除を実行
	return r.db.Delete(&models.Tag{}, tagID).Error
}
