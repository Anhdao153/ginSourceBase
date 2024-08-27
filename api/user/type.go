package user

type UserRequest struct {
}

type UserResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
