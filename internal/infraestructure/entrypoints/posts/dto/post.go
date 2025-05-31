package dto

type CreatePostRequestDTO struct {
	Title   string   `json:"title" validator:"required,min=3,max=100"`
	Content string   `json:"content" validator:"required,min=3,max=1000"`
	Tags    []string `json:"tags"`
	UserID  int      `json:"user_id"`
}

type PostResponseDTO struct {
	ID int `json:"id"`
	*CreatePostRequestDTO
}
