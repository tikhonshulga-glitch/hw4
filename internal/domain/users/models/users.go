package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Usertoken struct {
	AccessToken string `json:"asess_token"`
}
