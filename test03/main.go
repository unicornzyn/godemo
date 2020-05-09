package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 随机生成1-100的数字 直到生成99 打印随机的次数
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		fmt.Println(num)
		count++
		if num == 99 {
			fmt.Printf("生成99一共用了%d次", count)
			break
		}
	}
}
