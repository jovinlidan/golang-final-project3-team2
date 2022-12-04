package task_resources

type TaskCreateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryId  int64  `json:"category_id" validate:"required"`
}

type TaskUpdateRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type TaskUpdateStatusRequest struct {
	Status bool `json:"status" validate:"required"`
}

type TaskUpdateCategoryRequest struct {
	CategoryId int64 `json:"category_id" validate:"required"`
}
