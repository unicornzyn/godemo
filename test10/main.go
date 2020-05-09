package main

import "fmt"

func main() {
	// 数组遍历
	arr := [...]string{"aaa", "bbb", "ccc", "ddd", "eee"}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		fmt.Printf("arr[%d]=%v\n", i, v)
	}
}
