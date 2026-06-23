package post

type CreatePostInput struct {
	Title   string `json:"title" binding:"required,min=5,max=100"`
	Content string `json:"content" binding:"required,min=5"`
}

type UpdatePostInput struct {
	Title   string `json:"title" binding:"omitempty,min=5,max=100"`
	Content string `json:"content" binding:"omitempty,min=5"`
}
