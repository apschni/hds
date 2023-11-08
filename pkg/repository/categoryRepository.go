package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
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

func (g *CategoryRepository) GetSubjectsAll(ctx context.Context) ([]model.Subject, error) {
	var subjects []model.Subject

	query := "SELECT * FROM subjects"

	err := g.db.SelectContext(ctx, &subjects, query)
	if err != nil {
		return nil, errors.NewInternal()
	}
	return subjects, nil
}
