package seeds

import (
	"learn-gin/internal/post"
	"log"

	"gorm.io/gorm"
)

func SeedPosts(db *gorm.DB) error {
	var count int64
	var posts []post.Post

	tableName := post.Post.TableName(post.Post{})
	countResult := db.Model(&posts).Count(&count)

	if countResult.Error != nil {
		return countResult.Error
	}

	if count >= 1 {
		log.Printf("Table %s has data already. Skipped...", tableName)
		return nil
	}

	posts = []post.Post{
		{Title: "Sample Post #1", Content: "This is a sample post."},
		{Title: "Sample Post #2", Content: "This is another sample post."},
		{Title: "Sample Post #3", Content: "This is a sample post with no content."},
		{Title: "Sample Post #4", Content: "This is a sample post with no content."},
		{Title: "Sample Post #5", Content: "This is a sample post with no content."},
	}

	createResult := db.Create(&posts)

	if createResult.Error != nil {
		return createResult.Error
	}

	log.Printf("Table %s seeded with %d data.", tableName, createResult.RowsAffected)

	return nil
}
