package post

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostService struct {
	db *gorm.DB
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{
		db: db,
	}
}

func (s *PostService) GetAllPosts() ([]Post, error) {
	var posts []Post

	if err := s.db.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetPostById(id int) (Post, error) {
	var post Post

	if err := s.db.First(&post, id).Error; err != nil {
		return Post{}, err
	}

	return post, nil
}

func (s *PostService) CreatePost(input CreatePostInput) (Post, error) {
	post := Post{
		Title:   input.Title,
		Content: input.Content,
	}

	if err := s.db.Create(&post).Error; err != nil {
		return Post{}, err
	}

	return post, nil
}

func (s *PostService) UpdatePost(id int, input UpdatePostInput) (Post, error) {
	var post Post

	result := s.db.Model(&post).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(input)

	if result.Error != nil {
		return Post{}, result.Error
	}

	if result.RowsAffected == 0 {
		return Post{}, gorm.ErrRecordNotFound
	}

	return post, nil
}

func (s *PostService) DeletePost(id int) error {
	var post Post

	result := s.db.Delete(&post, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
