package process

import (
	"fmt"

	"github.com/unicornzyn/gotest27/common/message"
	"github.com/unicornzyn/gotest27/model"
)

var onlineUsers map[int]*model.User = make(map[int]*model.User, 10)
var currUser CurrUser

func updateUserStatus(n *message.NotifyUserStatusMsg) {
	user, ok := onlineUsers[n.UserID]
	if ok {
		user.UserStatus = n.Status
		return
	}
	user = &model.User{
		UserID:     n.UserID,
		UserStatus: n.Status,
	}
	onlineUsers[n.UserID] = user
	outputOnlineUsers()
}

func outputOnlineUsers() {
	fmt.Println("当前在线用户列表:")
	for id, user := range onlineUsers {
		fmt.Printf("用户id:%d\t状态:%v\n", id, user.UserStatus)
	}
}
