package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/unicornzyn/gotest27/model"

	"github.com/unicornzyn/gotest27/common/message"
	"github.com/unicornzyn/gotest27/common/utils"
)

// UserProcessor struct
type UserProcessor struct {
}

// Login 登录验证
func (u *UserProcessor) Login(userID int, userPwd string) (err error) {
	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	// 延时关闭
	defer conn.Close()

	// 组织数据
	var msg message.Message
	msg.Type = message.LoginMsgType
	loginMsg := message.LoginMsg{
		UserID:  userID,
		UserPwd: userPwd,
	}
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("loginMsg json.Marshal err=", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg json.Marshal err=", err)
		return
	}

	// 发送数据给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)

	// 等待服务端返回数据
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) err=", err)
		return
	}

	var loginRes message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginRes)
	if err != nil {
		fmt.Println("json.Unmarshal(msg.Data, &loginRes) err=", err)
		return
	}

	if loginRes.Code == 200 {
		fmt.Println("登录成功")
		// 初始化当前用户
		currUser.Conn = conn
		currUser.UserID = userID
		currUser.UserStatus = message.UserOnline

		// 初始化在线用户
		for _, id := range loginRes.UserIds {
			user := &model.User{
				UserID:     id,
				UserStatus: message.UserOnline,
			}
			onlineUsers[id] = user
		}
		fmt.Print("\n\n")
		go serverProcessMes(conn)
		ShowMenu()
	} else {
		fmt.Println(loginRes.Error)
	}

	return
}

// Register 注册
func (u *UserProcessor) Register(user *model.User) (err error) {
	// 连接到服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	// 延时关闭
	defer conn.Close()

	// 组织数据
	var msg message.Message
	msg.Type = message.RegisterMsgType
	registerMsg := message.RegisterMsg{
		User: *user,
	}
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("registerMsg json.Marshal err=", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg json.Marshal err=", err)
		return
	}

	// 发送数据给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)

	// 等待服务端返回数据
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) err=", err)
		return
	}

	var registerRes message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerRes)
	if err != nil {
		fmt.Println("json.Unmarshal(msg.Data, &registerRes) err=", err)
		return
	}

	if registerRes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
	} else {
		fmt.Println(registerRes.Error)
	}

	return
}
