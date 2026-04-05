package post

// CreatePostInput digunakan untuk parsing dan validasi request body saat membuat post
type CreatePostInput struct {
	Title   string `json:"title" binding:"required,min=5,max=100"`
	Content string `json:"content" binding:"required,min=5"`
}

// UpdatePostInput (Opsional) jika Anda ingin field yang berbeda saat update
type UpdatePostInput struct {
	Title   string `json:"title" binding:"omitempty,min=5"`
	Content string `json:"content" binding:"omitempty"`
}
