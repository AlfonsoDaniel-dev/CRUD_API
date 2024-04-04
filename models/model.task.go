package models

import (
	"time"
)

type Task struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Created_at time.Time `json:"created_At"`
}

type Tasks []*Task
