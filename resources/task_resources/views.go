package task_resources

import "time"

type TaskCreateResponse struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      int64     `json:"user_id"`
	CategoryId  int64     `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskGetUserResponse struct {
	Id       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type TaskGetResponse struct {
	Id          int64                `json:"id"`
	Title       string               `json:"title"`
	Status      bool                 `json:"status"`
	Description string               `json:"description"`
	UserId      int64                `json:"user_id"`
	CategoryId  int64                `json:"category_id"`
	CreatedAt   time.Time            `json:"created_at"`
	User        *TaskGetUserResponse `json:"user"`
}

type TaskUpdateResponse struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Status      bool      `json:"status"`
	Description string    `json:"description"`
	UserId      int64     `json:"user_id"`
	CategoryId  int64     `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskUpdateStatusResponse struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int64     `json:"user_id"`
	CategoryId  int64     `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskUpdateCategoryResponse struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int64     `json:"user_id"`
	CategoryId  int64     `json:"category_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
