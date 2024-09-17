package user

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password bool   `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
