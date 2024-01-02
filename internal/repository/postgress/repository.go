package repository

import "github.com/jmoiron/sqlx"

type TaskRepository interface {
}

type Repository struct {
	TaskRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TaskRepository: NewTaskRepository(db),
	}
}
