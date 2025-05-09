package main

import (
	"log"
	"os"

	"reminder/src/routes"
	"reminder/src/service"
	"reminder/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

func main() {
	// if err := utils.LoadEnv(); err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	redisOpt := asynq.RedisClientOpt{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
		Username: os.Getenv("REDIS_USER"),
		DB:       0,
	}

	taskProcessor := service.NewTaskProcessor(redisOpt)
	go func() {
		if err := taskProcessor.Start(); err != nil {
			log.Fatalf("Failed to start Asynq worker: %v", err)
		}
	}()
	log.Println("Asynq worker started")

	r := gin.Default()
	r.Use(utils.SetupCORS())

	routes.SetUpAllRoutes(r, redisOpt)

	port := utils.GetPort()
	log.Printf("Starting server on port %s", port)
	r.Run(":" + port)
}
