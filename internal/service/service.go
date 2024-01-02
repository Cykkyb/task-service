package service

import (
	"log/slog"
	repository "task-service/internal/repository/postgress"
)

type Task interface {
}

type Service struct {
	Task
}

func NewService(repo repository.TaskRepository, log *slog.Logger) *Service {
	return &Service{
		Task: NewTaskService(repo, log),
	}
}
