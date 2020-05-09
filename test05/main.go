package main

import "fmt"

// AddUpper 累加器 匿名函数 闭包 示例
func AddUpper() func(int) int {
	var n int = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	fn := AddUpper()
	fmt.Println(fn(1))
	fmt.Println(fn(2))
	fmt.Println(fn(3))
}
