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
	var resps []dto.GetTaskResp

	tasks, err := t.taskRepo.GetByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		teacher, err := t.userRepo.FindById(ctx, task.TeacherId)
		if err != nil {
			return nil, err
		}
		resps = append(resps, dto.GetTaskResp{
			Id:        task.Id,
			Label:     task.Label,
			Text:      task.Text,
			Subject:   task.Subject,
			Teacher:   teacher.FullName,
			Keypoint:  task.IsKeyPoint,
			Points:    task.Points,
			Completed: task.Closed,
			Deadline:  task.Deadline,
		})
	}

	return resps, nil
}
