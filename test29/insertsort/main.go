package main

import (
	"fmt"
)

// insertSort 插入排序 倒序
func insertSort(arr *[5]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
	}

}

func main() {
	arr := [5]int{23, 0, 52, 16, 34}
	insertSort(&arr)
	fmt.Println(arr)
}
