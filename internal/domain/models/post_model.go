package domain

type PostModel struct {
	ID        int
	Title     string
	Content   string
	Tags      []string
	UserID    int
	CreatedAt string
	UpdatedAt string
}
