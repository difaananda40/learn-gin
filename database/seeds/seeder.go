package seeds

import (
	"log"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	log.Println("Running database seeding...")

	seeders := []func(*gorm.DB) error{
		SeedPosts,
	}

	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			log.Fatalf("Failed to run seeder: %v", err)
		}
	}

	log.Println("Database seeding completed successfully!")
}
