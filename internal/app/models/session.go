package models

import "time"

type LoginSessionInfo struct {
	Username  string    `json:"username"`
	Token     string    `json:"token"`
	LoginTime time.Time `json:"login_time"`
}

func (u *LoginSessionInfo) TableName() string {
	return "login_session_info"
}
