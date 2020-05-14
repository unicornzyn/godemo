package main

import (
	"fmt"
)

func main() {
	// 统计1-8000中的素数 1个协程写 4个协程读 添加了 channel 只读(<-chan) 只写(chan<-) 的实践
	num := 4                          //计算协程数 一般等于逻辑cpu数量即可
	intChan := make(chan int, 1000)   //输入数
	primeChan := make(chan int, 1000) //素数
	exitChan := make(chan bool, num)  //退出标示

	//写入待计算的数
	go func(intChan chan<- int) {
		for i := 2; i <= 8000; i++ {
			intChan <- i
		}
		close(intChan)
	}(intChan)

	//判断每个数是否是素数
	for i := 0; i < num; i++ {
		go func(intChan <-chan int, primeChan chan<- int, exitChan chan<- bool, i int) {
			for {
				flag := true
				v, ok := <-intChan
				if !ok {
					break
				}
				for k := 2; k < v; k++ {
					if v%k == 0 {
						flag = false
						break
					}
				}
				if flag {
					primeChan <- v
				}
			}
			exitChan <- true
			fmt.Println("协程结束~ ", i)
		}(intChan, primeChan, exitChan, i)
	}

	//全部计算关闭PrimeChan
	go func(exitChan <-chan bool, primeChan chan<- int, num int) {
		for i := 0; i < num; i++ {
			<-exitChan
		}
		close(primeChan)
	}(exitChan, primeChan, num)

	//输出结果
	ok := true
	v := 0
	for ok {
		select {
		case v, ok = <-primeChan:
			if ok {
				fmt.Println("素数 =", v)
			}
		}

	}
}
