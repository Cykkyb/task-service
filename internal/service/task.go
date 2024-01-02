package service

import (
	"log/slog"
	repository "task-service/internal/repository/postgress"
)

type TaskService struct {
	log        *slog.Logger
	repository repository.TaskRepository
}

func NewTaskService(repository repository.TaskRepository, log *slog.Logger) *TaskService {
	return &TaskService{
		log:        log,
		repository: repository,
	}
}
