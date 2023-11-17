package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"homeworkdeliverysystem/dto"
	"homeworkdeliverysystem/errors"
	"homeworkdeliverysystem/model"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (g *CategoryRepository) GetCategoriesAll(ctx context.Context) ([]model.Category, error) {
	var subjects []model.Category

	query := "SELECT * FROM categories"

	err := g.db.SelectContext(ctx, &subjects, query)
	if err != nil {
		return nil, errors.NewInternal()
	}
	return subjects, nil
}

func (g *CategoryRepository) GetSubjectsAll(ctx context.Context, category_id uuid.UUID) (dto.GetSubjFromCategory, error) {
	var subjects dto.GetSubjFromCategory

	query := "SELECT subjects FROM categories where categories.id = $1"

	err := g.db.GetContext(ctx, &subjects, query, category_id)
	if err != nil {
		return subjects, errors.NewInternal()
	}
	return subjects, nil
}
