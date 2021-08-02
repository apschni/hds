package service

import (
	"context"
	"github.com/google/uuid"
	"homeworkdeliverysystem/dto"
	"homeworkdeliverysystem/model"
	"homeworkdeliverysystem/pkg/repository"
)

type TaskService struct {
	taskRepo repository.Task
	userRepo repository.User
}

func NewTaskService(taskRepo repository.Task, userRepo repository.User) *TaskService {
	return &TaskService{taskRepo: taskRepo, userRepo: userRepo}
}

func (t *TaskService) Create(ctx context.Context, task *model.Task) (string, error) {
	id, err := t.taskRepo.Create(ctx, *task)
	return id, err
}

func (t *TaskService) GetByUserId(ctx context.Context, id uuid.UUID) ([]dto.GetTaskResp, error) {
	tasks, err := t.taskRepo.GetByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
