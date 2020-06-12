package main

import (
	"fmt"
)

// selectsort 选择排序  倒序
func selectsort(arr *[5]int) {
	for j := 0; j < len(arr); j++ {
		max := arr[j]
		maxIndex := j
		// 找到最大值
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		// 交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}
	}

}

func main() {
	arr := [5]int{10, 30, 19, 100, 80}
	selectsort(&arr)
	fmt.Println(arr)
}
