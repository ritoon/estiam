package model

// go to Model folder
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}
