package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/unicornzyn/gotest27/common/message"

	"github.com/unicornzyn/gotest27/common/utils"
)

// ShowMenu 显示登录后菜单
func ShowMenu() {
	for {
		fmt.Println("---------恭喜xxx登录成功---------")
		fmt.Println("---------1. 显示在线用户列表---------")
		fmt.Println("---------2. 发送消息---------")
		fmt.Println("---------3. 信息列表---------")
		fmt.Println("---------4. 退出系统---------")
		fmt.Println("请选择(1-4):")
		var key int
		fmt.Scanf("%d\n", &key)
		var content string
		sms := &SmsProcessor{}
		switch key {
		case 1:
			outputOnlineUsers()
		case 2:
			fmt.Println("请输入群发消息内容:")
			fmt.Scanf("%s\n", &content)
			sms.SendGroupMsg(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择了退出系统")
			os.Exit(0)
		default:
			fmt.Println("选择错误")
		}
	}
}

// serverProcessMes 和服务器保持连接 接收服务器端消息
func serverProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}

	for {
		fmt.Println("客户端等待服务器端的消息...")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err=", err)
			return
		}
		//fmt.Println(mes)
		switch msg.Type {
		case message.NotifyUserStatusMsgType: // 用户上线通知
			var n message.NotifyUserStatusMsg
			fmt.Println(msg.Data)
			err = json.Unmarshal([]byte(msg.Data), &n)
			if err != nil {
				fmt.Println("json.Unmarshal(msg.Data) err=", err)
				return
			}
			updateUserStatus(&n)
		case message.SmsMsgType: //群发消息通知
			var smsMsg message.SmsMsg
			err = json.Unmarshal([]byte(msg.Data), &smsMsg)
			if err != nil {
				fmt.Println("json.Unmarshal([]byte(msg.Data), &smsMsg) err=", err)
				return
			}
			fmt.Printf("用户id:%d 对大家说:%s\n\n", smsMsg.User.UserID, smsMsg.Content)
		default:
			fmt.Println("服务器返回未知消息类型")
		}

	}

}
