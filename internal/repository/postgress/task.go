package repository

import "github.com/jmoiron/sqlx"

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{
		db: db,
	}
}
