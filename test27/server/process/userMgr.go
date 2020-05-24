package process

import (
	"fmt"
)

// UserMgr 管理在线用户列表
type UserMgr struct {
	onlineUsers map[int]*UserProcessor
}

// 全局变量
var userMgr *UserMgr

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcessor, 1024),
	}
}

// AddOnlineUser 添加/修改在线用户
func (m *UserMgr) AddOnlineUser(up *UserProcessor) {
	m.onlineUsers[up.UserID] = up
}

// DeleteOnlineUser 删除在线用户
func (m *UserMgr) DeleteOnlineUser(userID int) {
	delete(m.onlineUsers, userID)
}

// GetAllOnlineUsers 在线用户列表
func (m *UserMgr) GetAllOnlineUsers() map[int]*UserProcessor {
	return m.onlineUsers
}

// GetOnlineUserByID 返回id对应的值
func (m *UserMgr) GetOnlineUserByID(userID int) (up *UserProcessor, err error) {
	up, ok := m.onlineUsers[userID]
	if !ok {
		err = fmt.Errorf("用户%d不在线", userID)
	}
	return
}
