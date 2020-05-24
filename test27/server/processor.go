package server

import (
	"fmt"
	"io"
	"net"

	"github.com/unicornzyn/gotest27/common/utils"

	"github.com/unicornzyn/gotest27/server/process"

	"github.com/unicornzyn/gotest27/common/message"
)

// Processor struct
type Processor struct {
	Conn net.Conn
}

// ServerProcessMsg 根据不同类型处理消息
func (p *Processor) serverProcessMsg(msg *message.Message) (err error) {
	switch msg.Type {
	case message.LoginMsgType: //登录
		u := &process.UserProcessor{
			Conn: p.Conn,
		}
		err = u.Login(msg)
	case message.RegisterMsgType: //注册
		u := &process.UserProcessor{
			Conn: p.Conn,
		}
		err = u.Register(msg)
	case message.SmsMsgType: //群发消息
		smsProcessor := &process.SmsProcessor{}
		smsProcessor.SendGroupMsg(msg)
	default:
		fmt.Println("消息类型不存在")
	}
	return
}

// Process method
func (p *Processor) Process() (err error) {
	for {
		tf := &utils.Transfer{
			Conn: p.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端连接关闭")
			} else {
				fmt.Println("utils.ReadPkg(conn) err=", err)
			}
			return err
		}
		err = p.serverProcessMsg(&msg)
	}
}
