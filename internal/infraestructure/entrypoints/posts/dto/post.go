package dto

type CreatePostRequestDTO struct {
	Title   string   `json:"title" validate:"required,min=3,max=100" example:"My first post"`
	Content string   `json:"content" validate:"required,min=3,max=1000" example:"This is the content of my first post"`
	Tags    []string `json:"tags" example:"first,post"`
	UserID  int      `json:"user_id"`
}

type PostWithMetadata struct {
	ID int `json:"id"`
	*CreatePostRequestDTO
}

type PostSuccessAPIResponse struct {
	Data *PostWithMetadata `json:"data"`
}
