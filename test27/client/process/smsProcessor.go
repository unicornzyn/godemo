package process

import (
	"encoding/json"
	"fmt"

	"github.com/unicornzyn/gotest27/common/utils"

	"github.com/unicornzyn/gotest27/common/message"
)

// SmsProcessor 消息结构体
type SmsProcessor struct {
}

// SendGroupMsg 群发信息
func (p *SmsProcessor) SendGroupMsg(content string) (err error) {
	var msg message.Message
	msg.Type = message.SmsMsgType

	smsMsg := message.SmsMsg{
		Content: content,
		User:    currUser.User,
	}

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("json.Marshal(smsMsg) err=", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal(msg) err=", err)
		return
	}
	tf := utils.Transfer{
		Conn: currUser.Conn,
	}

	err = tf.WritePkg(data)
	return
}
