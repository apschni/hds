package handler

import (
	"github.com/gin-gonic/gin"
	apperrors "homeworkdeliverysystem/errors"
	"homeworkdeliverysystem/pkg/handler/middleware"
	"homeworkdeliverysystem/pkg/service"
	"os"
	"strconv"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	app := router.Group("/")

	handlerTimeout := os.Getenv("HANDLER_TIMEOUT")
	ht, _ := strconv.ParseInt(handlerTimeout, 0, 64)

	app.Use(middleware.Timeout(time.Duration(ht)*time.Second, apperrors.NewServiceUnavailable()))

	auth := app.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)                                          //registration
		auth.POST("/sign-in", h.SignIn)                                          //authentication
		auth.POST("/tokens", h.Tokens)                                           //refresh token
		auth.POST("/sign-out", middleware.AuthUser(h.services.Token), h.SignOut) //signing out
	}

	api := app.Group("/api", middleware.AuthUser(h.services.Token))
	{
		user := api.Group("/user")
		{
			user.GET("/me", h.Me)             //get user that currently logged in
			user.GET("/tasks", h.getAllTasks) //get all tasks ordered by deadline
		}

		tasks := api.Group("/tasks")
		{
			tasks.POST("/", middleware.Authority("teacher admin"), h.createTask) //create task
			tasks.PATCH("/", h.updateWithFile)                                   //прикрепление файла к таске
			tasks.POST("/:id/answer", h.answerTask)                              //прикрепить ответ на таску

			task := api.Group("/:id")
			{
				task.POST("/open", h.openTask)        //открыть таску
				task.POST("/close", h.closeTask)      //закрыть таску
				task.POST("/approve", h.approveTask)  //аппрувнуть ответ и закрыть таску
				task.POST("/rate", h.rateTask)        //оценить таску поинтами
				task.GET("/answers", h.getAllAnswers) //получить все ответы на таску (в порядке их создания)
			}
		}
	}

	return router
}
