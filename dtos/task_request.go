package dtos

type CreateTaskRequest struct {
	Title string  `json:"title" validate:"required"`
	Body  *string `json:"body,omitempty"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title,omitempty"`
	Body        *string `json:"body,omitempty"`
	IsCompleted *bool   `json:"is_completed,omitempty"`
}