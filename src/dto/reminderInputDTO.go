package dto

import (
	"time"
)

type ReminderInputDTO struct {
	To      string    `json:"to" binding:"required,email"`
	Subject string    `json:"subject" binding:"required"`
	Body    string    `json:"body" binding:"required"`
	SendAt  time.Time `json:"sendAt" binding:"required"`
}
