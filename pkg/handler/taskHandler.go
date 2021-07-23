package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"homeworkdeliverysystem/dto"
	apperrors "homeworkdeliverysystem/errors"
	"homeworkdeliverysystem/model"
	"log"
	"net/http"
	"time"
)

func (h *Handler) createTask(ctx *gin.Context) {
	userFromContext, _ := ctx.Get("user")
	user := userFromContext.(*model.User)

	var req dto.CreateTaskReq

	if ok := bindData(ctx, &req); !ok {
		return
	}

	now := time.Now()

	task := &model.Task{
		Label:      req.Label,
		Subject:    req.Subject,
		Text:       req.Text,
		Deadline:   req.Deadline,
		Points:     req.Points,
		Closed:     false,
		TeacherId:  user.Id,
		StudentId:  req.StudentId,
		FileName:   req.FileName,
		CreatedAt:  now,
		UpdatedAt:  now,
		IsKeyPoint: req.IsKeyPoint,
	}
	c := ctx.Request.Context()

	id, err := h.services.Task.Create(c, task)
	if err != nil {
		log.Printf("Failed create task: %v\n", err.Error())
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	taskId, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"task_id": taskId,
	})
}

func (h *Handler) getAllTasks(ctx *gin.Context) {
	userFromContext, _ := ctx.Get("user")
	user := userFromContext.(*model.User)

	c := ctx.Request.Context()

	tasks, err := h.services.Task.GetByUserId(c, user.Id)
	if err != nil {
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

func (h *Handler) answerTask(ctx *gin.Context) {

}

func (h *Handler) updateWithFile(ctx *gin.Context) {

}

func (h *Handler) openTask(ctx *gin.Context) {

}

func (h *Handler) closeTask(ctx *gin.Context) {

}

func (h *Handler) approveTask(ctx *gin.Context) {

}

func (h *Handler) rateTask(ctx *gin.Context) {

}

func (h *Handler) getAllAnswers(ctx *gin.Context) {

}
