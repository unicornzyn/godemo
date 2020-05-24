package client

import (
	"fmt"
	"os"

	"github.com/unicornzyn/gotest27/model"

	"github.com/unicornzyn/gotest27/client/process"
)

var (
	user model.User
)

// Run client run
func Run() {
	// 接收用户的选择
	var key int
	for {
		fmt.Println("-------------------欢迎登录多人聊天系统-------------------")
		fmt.Println("\t\t\t1 登录聊天室")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("\t\t\t4 请选择(1-3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的id:")
			fmt.Scanf("%d\n", &user.UserID)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &user.UserPwd)
			up := &process.UserProcessor{}
			up.Login(user.UserID, user.UserPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的id:")
			fmt.Scanf("%d\n", &user.UserID)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &user.UserPwd)
			fmt.Println("请输入用户的姓名:")
			fmt.Scanf("%s\n", &user.UserName)
			up := &process.UserProcessor{}
			up.Register(&user)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)

		default:
			fmt.Println("你的输入有误，请重新输入 ")
		}
	}
}
