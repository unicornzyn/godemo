package main

import "fmt"

func main() {
	// 打印空心金字塔
	total := 10
	for i := 0; i < total; i++ {
		for j := 0; j < total-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			if j == 0 || j == 2*i-2 || i == total-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Print("\n")
	}
}
