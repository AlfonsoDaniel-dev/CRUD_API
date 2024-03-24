package models

import (
	"time"
)

type Task struct {
	ID         int
	Title      string
	Content    string
	Created_at time.Time
}
