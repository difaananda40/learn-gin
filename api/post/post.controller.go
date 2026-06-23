package post

import (
	"errors"
	"learn-gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandlers(r *gin.RouterGroup, db *gorm.DB) {
	ctrl := &postController{
		service: NewPostService(db),
	}

	routes := r.Group("/post")
	{
		routes.GET("", ctrl.Index)
		routes.GET("/:id", ctrl.Show)
		routes.POST("", ctrl.Create)
		routes.PATCH("/:id", ctrl.Update)
		routes.DELETE("/:id", ctrl.Delete)
	}
}

type postController struct {
	service *PostService
}

func (ctrl *postController) Index(c *gin.Context) {
	posts, err := ctrl.service.GetAllPosts()

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "There's an issue in server.", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully.", posts)
}

func (ctrl *postController) Show(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID.", nil)
		return
	}

	post, err := ctrl.service.GetPostById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(c, http.StatusNotFound, "Post cannot be found.", nil)
			return
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Post retrieved successfully.", post)
}

func (ctrl *postController) Create(c *gin.Context) {
	var input CreatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErr := utils.FormatValidationError(err)

		utils.ErrorResponse(c, http.StatusBadRequest, validationErr.Message, validationErr.Errors)
		return
	}

	post, err := ctrl.service.CreatePost(input)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Post failed to create.", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Post successfully created.", post)
}

func (ctrl *postController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID.", nil)
		return
	}

	var input UpdatePostInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErr := utils.FormatValidationError(err)

		utils.ErrorResponse(c, http.StatusBadRequest, validationErr.Message, validationErr.Errors)
		return
	}

	post, err := ctrl.service.UpdatePost(id, input)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(c, http.StatusNotFound, "Post cannot be found.", nil)
			return
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, "Post failed to update.", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Post successfully updated.", post)
}

func (ctrl *postController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid post ID.", nil)
		return
	}

	if err := ctrl.service.DeletePost(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ErrorResponse(c, http.StatusNotFound, "Post cannot be found.", nil)
			return
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, "Post failed to delete.", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Post successfully deleted.", nil)
}
