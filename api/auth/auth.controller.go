package auth

import (
	"net/http"

	"learn-gin/config"
	"learn-gin/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandlers(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	ctrl := &authController{
		service: NewAuthService(db, cfg),
	}

	routes := r.Group("/auth")
	{
		routes.POST("/login", ctrl.Login)
		routes.POST("/register", ctrl.Register)
	}
}

type authController struct {
	service *AuthService
}

func (ctrl *authController) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErr := utils.FormatValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, validationErr.Message, validationErr.Errors)
		return
	}

	user, token, err := ctrl.service.Login(input)

	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	userResponse := ToLoginResponse(user, token)

	utils.SuccessResponse(c, http.StatusOK, "Login successful", userResponse)
}

func (ctrl *authController) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErr := utils.FormatValidationError(err)
		utils.ErrorResponse(c, http.StatusBadRequest, validationErr.Message, validationErr.Errors)
		return
	}

	user, err := ctrl.service.Register(input)

	if err != nil {
		validationErr := utils.FormatValidationError(err)
		if validationErr.Message == "Invalid validation" {
			utils.ErrorResponse(c, http.StatusConflict, validationErr.Message, validationErr.Errors)
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to register user", nil)
		return
	}

	userResponse := ToUserResponse(user)
	utils.SuccessResponse(c, http.StatusOK, "User registered successfully", userResponse)
}
