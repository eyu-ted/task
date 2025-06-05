# Task API

This is a simple REST API built using Go and Gin. It stores tasks in memory.

## Endpoints

- GET /api/tasks - Get all tasks
- POST /api/tasks - Add a new task
- PUT /api/tasks/:id - Mark a task as completed
- DELETE /api/tasks/:id - Delete a task

## How to Run

bash
- cd backend
- go run main.go