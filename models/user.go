package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type SignUpInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Passowrd string `json:"password"`
}

type LoginIputo struct {
	Email    string `json:"email"`
	Passowrd string `json:"password"`
}
