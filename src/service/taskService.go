package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"reminder/src/entity"

	"github.com/hibiken/asynq"
)

const (
	TypeEmailReminder = "email:reminder"
)

type TaskProcessor struct {
	server       *asynq.Server
	emailService *EmailService
}

func NewTaskProcessor(redisOpt asynq.RedisClientOpt) *TaskProcessor {
	emailService := NewEmailService()
	return &TaskProcessor{
		server:       asynq.NewServer(redisOpt, asynq.Config{Concurrency: 10}),
		emailService: emailService,
	}
}

func (p *TaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEmailReminder, p.handleEmailReminderTask)

	log.Println("Starting Asynq task processor")
	return p.server.Start(mux)
}

func (p *TaskProcessor) handleEmailReminderTask(ctx context.Context, t *asynq.Task) error {
	var reminder entity.Reminder
	if err := json.Unmarshal(t.Payload(), &reminder); err != nil {
		return fmt.Errorf("failed to unmarshal reminder payload: %w", err)
	}

	log.Printf("Processing email reminder: %s to %s (IST: %s)",
		reminder.Subject,
		reminder.To,
		reminder.SendAt.Format(time.RFC3339))

	err := p.emailService.SendEmail(reminder.To, reminder.Subject, reminder.Body)
	if err != nil {
		return err
	}

	return nil
}

type TaskScheduler struct {
	client *asynq.Client
}

func NewTaskScheduler(redisOpt asynq.RedisClientOpt) *TaskScheduler {
	return &TaskScheduler{
		client: asynq.NewClient(redisOpt),
	}
}

func (s *TaskScheduler) ScheduleEmailReminder(reminder entity.Reminder) error {
	payload, err := json.Marshal(reminder)
	if err != nil {
		return fmt.Errorf("failed to marshal reminder: %w", err)
	}

	delay := time.Until(reminder.SendAt)
	if delay < 0 {
		return fmt.Errorf("cannot schedule a reminder in the past")
	}

	task := asynq.NewTask(TypeEmailReminder, payload)
	info, err := s.client.Enqueue(task, asynq.ProcessIn(delay))
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Printf("Enqueued task: id=%s queue=%s scheduled_at=%s (IST)",
		info.ID,
		info.Queue,
		reminder.SendAt.Format(time.RFC3339))
	return nil
}
