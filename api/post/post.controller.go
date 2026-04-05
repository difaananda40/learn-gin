package post

import (
	"net/http"

	"learn-gin/utils"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup) {
	ctrl := &postController{
		service: NewPostService(),
	}

	routes := r.Group("/post")
	{
		routes.GET("", ctrl.GetAll)
		routes.POST("", ctrl.Create)
	}
}

type postController struct {
	service *PostService
}

func (ctrl *postController) GetAll(c *gin.Context) {
	posts, err := ctrl.service.GetAllPosts()

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "There's an issue in server.", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Posts retrieved successfully.", posts)
}

func (ctrl *postController) Create(c *gin.Context) {
	var input CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatValidationError(err)

		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid validation", errors)
		return
	}

	post, err := ctrl.service.CreatePost(input)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Post failed to create.", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Post successfully created.", post)
}
