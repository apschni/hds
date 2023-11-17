package service

import (
	"context"
	"github.com/google/uuid"
	"homeworkdeliverysystem/dto"
	"homeworkdeliverysystem/model"
	"homeworkdeliverysystem/pkg/repository"
)

type CategoryService struct {
	categoryRepo repository.Category
}

func NewCategoryService(categoryRepo repository.Category) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (g *CategoryService) GetCategoriesS(ctx context.Context) ([]model.Category, error) {
	return g.categoryRepo.GetCategoriesAll(ctx)
}

func (g *CategoryService) GetSubjectsS(ctx context.Context, id uuid.UUID) (dto.GetSubjFromCategory, error) {
	return g.categoryRepo.GetSubjectsAll(ctx, id)
}
