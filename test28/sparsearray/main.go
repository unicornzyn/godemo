package main

import (
	"fmt"
)

type valNode struct {
	row int
	col int
	val int
}

func main() {
	// 稀疏数组
	// 原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //白子

	// 输出原始数组
	fmt.Println("原始数组")
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Printf("\n")
	}

	// 转成稀疏数组
	var sparseArr []valNode
	// 初始节点
	val := valNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, val)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				val = valNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, val)
			}
		}
	}

	//打印稀疏数组
	fmt.Println("稀疏数组")
	for i, v := range sparseArr {
		fmt.Printf("%d\t%d\t%d\t%d\n", i, v.row, v.col, v.val)
	}

	//稀疏数组反转回原始数组
	var chessMap2 [11][11]int
	for i, v := range sparseArr {
		if i == 0 {
			continue
		}
		chessMap2[v.row][v.col] = v.val
	}

	//打印恢复后的数组
	fmt.Println("恢复后的数组")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Printf("\n")
	}
}
