package auth

import (
	"errors"
	"learn-gin/config"
	"learn-gin/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db     *gorm.DB
	config *config.Config
}

func NewAuthService(db *gorm.DB, config *config.Config) *AuthService {
	return &AuthService{
		db:     db,
		config: config,
	}
}

func (s *AuthService) Login(input LoginInput) (User, string, error) {
	var user User

	if err := s.db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return User{}, "", errors.New("Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return User{}, "", errors.New("Invalid credentials")
	}

	tokenPayload := utils.JWTPayload{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	token, err := utils.GenerateToken(tokenPayload, s.config.JWTSecret)
	if err != nil {
		return User{}, "", errors.New("failed to generate token")
	}

	return user, token, nil
}

func (s *AuthService) Register(input RegisterInput) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Username: input.Username,
		Password: string(hashedPassword),
	}

	if err := s.db.Create(&user).Error; err != nil {
		return User{}, err
	}

	return user, nil
}
