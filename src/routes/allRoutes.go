package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"

	"reminder/src/controllers"
)

func SetUpAllRoutes(r *gin.Engine, redisOpt asynq.RedisClientOpt) {
	reminderController := controllers.NewReminderController(redisOpt)

	r.GET("/", reminderController.RenderHomePage)
	r.POST("/reminder", reminderController.CreateReminder)
}
