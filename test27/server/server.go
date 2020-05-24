package server

import (
	"fmt"
	"net"
	"time"

	"github.com/unicornzyn/gotest27/model"
)

func init() {
	//初始化redis连接池
	initPool("127.0.0.1:6379", 16, 0, time.Second*300)
	//初始化userdao
	model.MyUserDao = model.UserDaoNew(pool)
}

// Run server run
func Run() {
	fmt.Println("服务器端口开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	for {
		fmt.Println("等待客户端连接...")
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通信
		go func(conn net.Conn) {
			defer conn.Close()

			//读取客户端的消息
			p := &Processor{
				Conn: conn,
			}
			p.Process()
		}(conn)
	}
}
