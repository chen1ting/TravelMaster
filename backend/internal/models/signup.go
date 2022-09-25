package models

type SignupReq struct {
	Username       string   `json:"username"`
	HashedPassword string   `json:"hashed_password"`
	Email          string   `json:"email"`
	Interests      []string `json:"interests"`
}

type SignupResp struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	SessionToken string `json:"session_token"`
}
