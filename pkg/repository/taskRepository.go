package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	apperrors "homeworkdeliverysystem/errors"
	"homeworkdeliverysystem/model"
	"log"
)

type TaskRepository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) Create(ctx context.Context, task model.Task) (string, error) {
	var id uuid.UUID
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", apperrors.NewInternal()
	}

	query := "INSERT INTO tasks " +
		"(id, label, subject, text, deadline, points, closed, teacher_id, file_name, student_id, created_at, updated_at, is_key_point)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id"

	err = t.db.GetContext(ctx, &id, query, newUUID, task.Label, task.Subject, task.Text, task.Deadline, task.Points, task.Closed,
		task.TeacherId, task.FileName, task.StudentId, task.CreatedAt, task.UpdatedAt, task.IsKeyPoint)

	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a task with label: %v. Reason: %v\n", task.Label, err.Code.Name())
			return "", apperrors.NewConflict("taskLabel", task.Label)
		}

		log.Printf("Could not create a task with label: %v. Reason: %v\n", task.Label, err)
		return "", apperrors.NewInternal()
	}

	return id.String(), nil
}

func (t *TaskRepository) GetByUserId(ctx context.Context, id uuid.UUID) ([]model.Task, error) {
	var tasks []model.Task

	query := "SELECT * FROM tasks WHERE student_id=$1 ORDER BY deadline"

	err := t.db.SelectContext(ctx, &tasks, query, id)
	if err != nil {
		log.Printf("Could not select a task with student_id: %v. Reason: %v\n", id, err)
		return nil, apperrors.NewInternal()
	}

	return tasks, nil
}