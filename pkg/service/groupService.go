package service

import (
	"context"
	"homeworkdeliverysystem/pkg/repository"
)

type GroupService struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) *GroupService {
	return &GroupService{repo: repo}
}

func (g *GroupService) GetByNumber(ctx context.Context, number string) ([]string, error) {
	return g.repo.GetSubjectsByGroupNumber(ctx, number)
}
