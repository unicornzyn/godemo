package message

import (
	"github.com/unicornzyn/gotest27/model"
)

// 消息类型常量
const (
	// LoginMsgType 登录消息类型
	LoginMsgType = "LoginMsg"
	// LoginResMsgType 登录结果消息
	LoginResMsgType = "LoginResMsg"
	// RegisterMsgType 注册消息类型
	RegisterMsgType = "RegisterMsg"
	// RegisterResMsgType 注册消息结果类型
	RegisterResMsgType = "RegisterResMsg"
	// NotifyUserStatusMsgType 用户状态消息类型
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	// SmsMsgType 短信消息
	SmsMsgType = "SmsMsg"
)

// 用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// Message struct
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

// LoginMsg 登录消息
type LoginMsg struct {
	UserID   int    `json:"userid"`
	UserPwd  string `json:"userpwd"`
	UserName string `json:"username"`
}

// LoginResMsg 登录结果消息
type LoginResMsg struct {
	Code    int    `json:"code"`
	UserIds []int  `json:"userids"`
	Error   string `json:"error"`
}

// RegisterMsg 注册消息
type RegisterMsg struct {
	model.User
}

// RegisterResMsg 注册结果消息
type RegisterResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// NotifyUserStatusMsg 通知用户状态消息
type NotifyUserStatusMsg struct {
	UserID int `json:"userid"`
	Status int `json:"status"`
}

// SmsMsg 发送的消息
type SmsMsg struct {
	Content string     `json:"content"`
	User    model.User `json:"user"`
}
