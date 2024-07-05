package services

import (
	"errors"
	"go-memo-api/internal/models"
	"go-memo-api/internal/repositories"
)

type TagService interface {
	GetTagByID(tagID int) (models.Tag, error)
	DeleteTag(tagID int) error
	CheckTagNotDeleted(tagID int) (bool, error)
}

type tagService struct {
	tagRepository repositories.TagRepository
}

func NewTagService(tagRepo repositories.TagRepository) TagService {
	return &tagService{tagRepo}
}

func (s *tagService) GetTagByID(tagID int) (models.Tag, error) {
	return s.tagRepository.GetTagByID(tagID)
}

func (s *tagService) DeleteTag(tagID int) error {
	return s.tagRepository.DeleteTag(tagID)
}

func (s *tagService) CheckTagNotDeleted(tagID int) (bool, error) {
	tag, err := s.tagRepository.GetTagByID(tagID)
	if err != nil {
		return false, err
	}
	if tag.DeletedAt.Valid {
		return false, errors.New("tag is already deleted")
	}
	return true, nil
}
