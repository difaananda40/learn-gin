package post

import (
	"time"
)

// Service struct untuk menampung dependensi (misal: DB)
type PostService struct {
	// db *gorm.DB  <-- Nanti masukkan koneksi DB di sini
}

// NewPostService adalah constructor
func NewPostService() *PostService {
	return &PostService{}
}

// GetAllPosts menangani logika pengambilan semua post
func (s *PostService) GetAllPosts() ([]Post, error) {
	// Simulasi data dari Database
	posts := []Post{
		{ID: 1, Title: "Belajar Go Gin Framework", Content: "Go itu seru!", CreatedAt: time.Now()},
		{ID: 2, Title: "Tips Fullstack", Content: "Gunakan struktur modular.", CreatedAt: time.Now()},
	}

	return posts, nil
}

// CreatePost menangani logika pembuatan post baru
func (s *PostService) CreatePost(input CreatePostInput) (Post, error) {
	// Simulasi simpan ke DB
	newPost := Post{
		ID:        uint(time.Now().Unix()), // Dummy ID
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now(),
	}

	return newPost, nil
}
