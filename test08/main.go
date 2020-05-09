package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("当前时间: %v now type: %T\n", now, now)

	//格式化日期时间 Sprintf也可以 除了返回字符串而不是直接输出 其他和Printf一样
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), int(now.Month()), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), int(now.Month()), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	fmt.Println(dateStr)

	//第二种格式化时间的方式 2006 01 02 15 04 05 对应年月日时分秒 必须按这个数字写
	fmt.Println(now.Format("2006年01月02日 15:04:05"))

	//时间常量的使用
	//间隔100毫秒输出 1-5  不能用time.Second/10 或者 *0.1 必须是整数运算
	for i := 1; i <= 5; i++ {
		fmt.Print(i)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("\n")

	//Unix时间戳(自1970-01-01的秒数) UnixNano(纳秒数)
	fmt.Println("Unix时间戳:", now.Unix())
}
