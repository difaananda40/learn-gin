package main

import (
	"flag"
	"learn-gin/api/auth"
	"learn-gin/api/post"
	"learn-gin/config"
	"learn-gin/database/seeds"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db := config.InitDB(cfg.DatabaseURL)

	db.AutoMigrate(&auth.User{}, &post.Post{})

	// Database seeder
	shouldSeed := flag.Bool("seed", false, "Run database seeder")
	flag.Parse()
	if *shouldSeed {
		seeds.SeedAll(db)
		os.Exit(0)
	}

	// Initialize router
	router := gin.Default()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true
	baseRoute := router.Group("/api")

	// Register handlers
	auth.RegisterHandlers(baseRoute, db, cfg)
	post.RegisterHandlers(baseRoute, db)

	// Initialize Validation
	config.InitializeValidation()

	router.Run() // listens on 0.0.0.0:8080 by default
}
