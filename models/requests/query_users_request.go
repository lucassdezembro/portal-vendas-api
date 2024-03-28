package models

type QueryUsersRequest struct {
	Document string `json:"document"`
	Email    string `json:"email"`
}
