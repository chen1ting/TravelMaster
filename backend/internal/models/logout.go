package models

type LogoutReq struct {
	SessionToken string `json:"session_token"`
}
