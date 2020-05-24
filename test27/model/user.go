package model

// User 定义user结构体
type User struct {
	UserID     int    `json:"userid"`
	UserPwd    string `json:"userpwd"`
	UserName   string `json:"username"`
	UserStatus int    `json:"userstatus"`
}
