package models

type LoginReq struct {
	//assumptions: use username to login
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
}

type LoginResp struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`

	SessionToken string `json:"session_token"`
}

type LogoutReq struct {
	SessionToken string `json:"session_token"`
}

type SignupReq struct {
	Username       string                   `form:"username"`
	HashedPassword string                   `form:"hashed_password"`
	Email          string                   `form:"email"`
	Avatar         *multipart.FileHeader    `form:"avatar"`
	 // Interests      []string `json:"interests"` update later, not as part of initial sign up req
}

type SignupResp struct {
	UserId    int64  `form:"user_id"`
	Username  string `form:"username"`
	Email     string `form:"email"`
	Avatarurl string `form:"avatarurl"`
  
	SessionToken string `form:"session_token"`
 }


type ValidateTokenReq struct {
	SessionToken string `json:"session_token"`
}

type ValidateTokenResp struct {
	Valid bool `json:"valid"`
	UserId int64 `json:"user_id"`
}