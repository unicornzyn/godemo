package main

import (
	"fmt"
)

// quickSort 快速排序 升序
func quickSort(left int, right int, arr *[6]int) {
	l := left
	r := right
	// 确定中轴
	pivot := arr[(left+right)/2]
	// 确定中轴左边全部是比中轴小的数 中轴右边全部是比中轴大的数
	for l < r {
		for arr[l] < pivot {
			l++
		}
		for arr[r] > pivot {
			r--
		}
		if l >= r {
			break
		}
		// 找到中轴 左边 大于中轴的数 和 中轴右边 小于中轴的数 交换位置
		arr[l], arr[r] = arr[r], arr[l]

		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		quickSort(left, r, arr)
	}
	if right > l {
		quickSort(l, right, arr)
	}

}

func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	quickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
