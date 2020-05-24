package process

import (
	"net"

	"github.com/unicornzyn/gotest27/model"
)

// CurrUser 当前登录用户
type CurrUser struct {
	Conn net.Conn
	model.User
}
