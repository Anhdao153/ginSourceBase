package user

type UserCreateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
	// RoleName string `json:"roleName" binding:"required"`
}

type UserUpdateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
	// RoleName string `json:"roleName" binding:"required"`
}

type UserDetailRequest struct {
	Id int `uri:"id" binding:"required"`
	// RoleName string `json:"roleName" binding:"required"`
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	RoleName string `json:"roleName"`
}
