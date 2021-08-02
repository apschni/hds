package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"homeworkdeliverysystem/dto"
	apperrors "homeworkdeliverysystem/errors"
	"homeworkdeliverysystem/model"
	"log"
	"net/http"
	"path"
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
		FileName:   "",
		CreatedAt:  now,
		UpdatedAt:  now,
		IsKeyPoint: req.IsKeyPoint,
	}
	c := ctx.Request.Context()

	id, err := h.services.Task.Create(c, task)
	if err != nil {
		log.Printf("Failed to create task: %v\n", err.Error())
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

func (h *Handler) GetAllTasks(ctx *gin.Context) {
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

func (h *Handler) UpdateTaskWithFile(ctx *gin.Context) {
	req := &dto.UploadFileOnTaskReq{}

	formFile, err := ctx.FormFile("file")
	if err != nil {
		log.Printf("Failed to get file from form: %v\n", err.Error())
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	req.File = formFile

	err = ctx.ShouldBindUri(&req)
	if err != nil {
		log.Printf("Failed to bind uri path params: %v\n", err.Error())
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	fileId, _ := uuid.NewRandom()

	req.File.Filename = fileId.String() + "_" + req.File.Filename

	c := ctx.Request.Context()

	err = h.services.Task.UpdateWithFile(c, req)
	if err != nil {
		log.Printf("Failed to update database with new filename: %v\n", err.Error())
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	dst := path.Join("./files/", req.File.Filename)
	err = ctx.SaveUploadedFile(req.File, dst)
	if err != nil {
		log.Printf("Failed save file: %v\n", err.Error())
		ctx.JSON(apperrors.Status(err), gin.H{
			"error": err,
		})
		return
	}
	ctx.String(http.StatusOK, "File uploaded successfully")
}
