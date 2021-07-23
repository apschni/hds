package repository

import (
	"context"
	"github.com/google/uuid"
	"homeworkdeliverysystem/model"
	"time"
)

type Task interface {
	Create(ctx context.Context, task model.Task) (string, error)
	GetByUserId(ctx context.Context, id uuid.UUID) ([]model.Task, error)
}

type User interface {
	Create(ctx context.Context, user model.User) (string, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	FindById(ctx context.Context, id uuid.UUID) (*model.User, error)
}

type Token interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
	DeleteUserRefreshToken(ctx context.Context, userID string) error
}

type Repository struct {
	Task
	User
	Token
}

func NewRepository(dataSources *dataSources) *Repository {
	return &Repository{
		User:  NewUserRepository(dataSources.DB),
		Token: NewTokenRepository(dataSources.RedisClient),
		Task:  NewTaskRepository(dataSources.DB),
	}
}
