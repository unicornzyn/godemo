package main

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(arr *[5]int) {
	var tmp int
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}

	}
}

func main() {
	arr := [5]int{24, 69, 80, 57, 13}
	BubbleSort(&arr)
	fmt.Println(arr)
}
