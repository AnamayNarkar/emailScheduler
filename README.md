# Email Reminder Service

A Go-based service for scheduling and sending email reminders at specified times. The application allows users to create reminders with customized subject, body, and delivery time.

## Features

- Schedule email reminders to be sent at a specific time
- Asynchronous task processing using Redis and Asynq
- RESTful API for creating reminders
- Web interface for creating reminders
- Time zone handling (IST/Asia Kolkata)

## Tech Stack

- **Backend**: Go
- **Web Framework**: Gin Gonic
- **Task Processing**: Asynq
- **Data Storage**: Redis
- **Email**: gopkg.in/mail.v2 (Gomail)

## Prerequisites

- Go 1.16+
- Redis server
- SMTP server for sending emails

## Environment Variables

Create a `.env` file in the root directory with the following variables:

```
# Server Configuration
PORT=8080

# Redis Configuration
REDIS_URL=localhost:6379
REDIS_PASS=

# Email Configuration
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=your-email@example.com
SMTP_PASS=your-password
EMAIL_FROM=your-email@example.com
```

## Installation & Setup

1. Clone the repository
2. Install dependencies:
   ```
   go mod tidy && go mod vendor
   ```
3. Set up your environment variables (see above)
4. Start Redis server

## Running the Application

### Development

```bash
go run main.go
```

### Production

```bash
./run.sh
```

## API Endpoints

### Create Reminder
- **URL**: `/reminder`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "to": "recipient@example.com",
    "subject": "Reminder Subject",
    "body": "This is your reminder message",
    "sendAt": "2025-05-10T15:00:00Z"
  }
  ```
- **Response**:
  ```json
  {
    "message": "Reminder scheduled successfully",
    "id": "unique-reminder-id",
    "sendAt": "2025-05-10T15:00:00Z"
  }
  ```

### Home Page
- **URL**: `/`
- **Method**: `GET`
- **Response**: HTML page for API documentation
