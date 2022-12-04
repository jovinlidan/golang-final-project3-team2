package task_domain

import "time"

type Task struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int64     `json:"user_id"`
	CategoryId  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
