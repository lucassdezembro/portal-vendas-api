package models

type CreateUserRequest struct {
	User CreateUserRequest_User `json:"user"`
}

type CreateUserRequest_User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
