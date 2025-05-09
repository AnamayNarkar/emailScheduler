package controllers

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"

	"reminder/src/dto"
	"reminder/src/entity"
	"reminder/src/service"
)

type ReminderController struct {
	taskScheduler *service.TaskScheduler
}

func NewReminderController(redisOpt asynq.RedisClientOpt) *ReminderController {
	return &ReminderController{
		taskScheduler: service.NewTaskScheduler(redisOpt),
	}
}

func (c *ReminderController) CreateReminder(ctx *gin.Context) {
	var input dto.ReminderInputDTO
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reminder := entity.Reminder{
		ID:      uuid.New().String(),
		To:      input.To,
		Subject: input.Subject,
		Body:    input.Body,
		SendAt:  input.SendAt,
	}

	err := c.taskScheduler.ScheduleEmailReminder(reminder)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Reminder scheduled successfully",
		"id":      reminder.ID,
		"sendAt":  reminder.SendAt.String(),
	})
}

func (c *ReminderController) RenderHomePage(ctx *gin.Context) {
	htmlFile := filepath.Join("src", "templates", "index.html")
	htmlContent, err := ioutil.ReadFile(htmlFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load template"})
		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", htmlContent)
}
