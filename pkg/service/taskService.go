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

func (t *TaskService) GetByUserId(ctx context.Context, category_id string, subjects_ids []string) ([]dto.GetTaskResp, error) {
	tasks, err := t.taskRepo.GetByUserId(ctx, category_id, subjects_ids)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskService) UpdateMultipleWithFile(ctx context.Context, req *dto.UpdateMultipleWithFileReq) error {
	return t.taskRepo.UpdateFileNameOnMultipleTasks(ctx, req.Ids, req.File.Filename)
}

func (t *TaskService) GetFileNameById(ctx context.Context, id string) (string, error) {
	fileName, err := t.taskRepo.GetFileNameById(ctx, id)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (t *TaskService) Open(ctx context.Context, id uuid.UUID) error {
	return t.taskRepo.Open(ctx, id)
}

func (t *TaskService) Close(ctx context.Context, id uuid.UUID) error {
	return t.taskRepo.Close(ctx, id)
}

func (t *TaskService) CheckAnswer(ctx context.Context, id uuid.UUID, answer string) (bool, string, error) {
	return t.taskRepo.CheckAnswer(ctx, id, answer)
}
