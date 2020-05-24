package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/unicornzyn/gotest27/model"

	"github.com/unicornzyn/gotest27/common/message"
	"github.com/unicornzyn/gotest27/common/utils"
)

// UserProcessor 用户处理
type UserProcessor struct {
	Conn   net.Conn
	UserID int
}

// Login 登录
func (p *UserProcessor) Login(msg *message.Message) (err error) {
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg.Data)) err=", err)
		return
	}
	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType
	var loginMsgRes message.LoginResMsg
	user, err := model.MyUserDao.Login(loginMsg.UserID, loginMsg.UserPwd)

	if err != nil {
		loginMsgRes.Code = 500
		loginMsgRes.Error = err.Error()
	} else {
		loginMsgRes.Code = 200
		//存入用户在线列表
		p.UserID = user.UserID
		userMgr.AddOnlineUser(p)
		for id := range userMgr.GetAllOnlineUsers() {
			loginMsgRes.UserIds = append(loginMsgRes.UserIds, id)
		}
		//通知其他用户当前用户上线了
		p.notifyOtherOnlineUsers()
	}
	data, err := json.Marshal(loginMsgRes)
	if err != nil {
		fmt.Println("json.Marshal(loginMsgRes) err=", err)
		return
	}
	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal(resMsg) err=", err)
		return
	}
	ft := &utils.Transfer{
		Conn: p.Conn,
	}
	err = ft.WritePkg(data)
	return
}

// Register 注册
func (p *UserProcessor) Register(msg *message.Message) (err error) {
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg.Data)) err=", err)
		return
	}
	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType
	var registerMsgRes message.RegisterResMsg
	err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		registerMsgRes.Code = 500
		registerMsgRes.Error = err.Error()
	} else {
		registerMsgRes.Code = 200
	}
	data, err := json.Marshal(registerMsgRes)
	if err != nil {
		fmt.Println("json.Marshal(registerMsgRes) err=", err)
		return
	}
	resMsg.Data = string(data)

	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal(resMsg) err=", err)
		return
	}
	ft := &utils.Transfer{
		Conn: p.Conn,
	}
	err = ft.WritePkg(data)
	return
}

// notifyOtherOnlineUsers 通知其他人当前用户上线
func (p *UserProcessor) notifyOtherOnlineUsers() {

	for id, up := range userMgr.onlineUsers {
		if id == p.UserID {
			continue
		}
		up.notifyOne(p.UserID)
	}
}

// notifyOne 通知指定的一个人
func (p *UserProcessor) notifyOne(userID int) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType
	notifyUserStatusMsg := message.NotifyUserStatusMsg{
		UserID: userID,
		Status: message.UserOnline,
	}
	data, err := json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("json.Marshal(notifyUserStatusMsg) err=", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)

	if err != nil {
		fmt.Println("json.Marshal(msg) err=", err)
		return
	}

	tf := utils.Transfer{
		Conn: p.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("tf.WritePkg(data) err=", err)
	}
}
