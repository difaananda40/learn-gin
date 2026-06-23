package auth

type LoginInput struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Password string `json:"password" binding:"required,min=5,max=120"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required,min=5,max=32"`
	Password string `json:"password" binding:"required,min=5,max=120"`
}
