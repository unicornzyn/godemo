package main

import "fmt"

// sum 可变参数测试
// 返回值可以命名参数 函数最终直接return即可 不用指定具体返回值
func sum(a int, args ...int) (v int) {
	v = a
	for i := 0; i < len(args); i++ {
		v += args[i]
	}
	return
}

func main() {
	fmt.Printf("10 + 20 + 30 = %d\n", sum(10, 20, 30))
}
