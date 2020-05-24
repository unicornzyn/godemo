package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/unicornzyn/gotest27/common/utils"

	"github.com/unicornzyn/gotest27/common/message"
)

// SmsProcessor 消息处理结构体
type SmsProcessor struct {
}

// SendGroupMsg 转发消息到客户端
func (p *SmsProcessor) SendGroupMsg(msg *message.Message) {
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(msg), &smsMsg) err=", err)
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal(msg) err=", err)
		return
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMsg.User.UserID {
			continue
		}
		p.sendMsgTo(data, up.Conn)
	}
}

// SendMsgTo 发给一个用户
func (p *SmsProcessor) sendMsgTo(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
}
