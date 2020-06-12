package main

import (
	"fmt"
)

// setWay 寻路(1,1)->(6,5) mymap是同一个地图的引用传递  i,j 表示地图的点坐标
func setWay(mymap *[8][7]int, i int, j int) bool {
	if mymap[6][5] == 2 {
		return true
	}
	if mymap[i][j] == 0 {
		// 假设这个点可以通 但是需要探测 探测方向顺序 下右上左
		mymap[i][j] = 2
		if setWay(mymap, i+1, j) { //下
			return true
		} else if setWay(mymap, i, j+1) { //右
			return true
		} else if setWay(mymap, i-1, j) { //上
			return true
		} else if setWay(mymap, i, j-1) { //左
			return true
		} else { // 当前点上下左右都不通 是死路
			mymap[i][j] = 3
			return false
		}
	} else {
		return false
	}
}

func main() {
	// 模拟迷宫
	// 规则
	// 1. 如果元素值为1 是墙
	// 2. 如果元素值为0 是没有走过的点
	// 3. 如果元素值为2 是一条通路
	// 4. 如果元素值为3 是走过的点 但是走不通

	// 初始化地图
	var mymap [8][7]int
	for i := 0; i < 7; i++ {
		mymap[0][i] = 1
		mymap[7][i] = 1
	}
	for i := 0; i < 8; i++ {
		mymap[i][0] = 1
		mymap[i][6] = 1
	}
	mymap[3][1] = 1
	mymap[3][2] = 1

	// 寻路
	setWay(&mymap, 1, 1)

	// 输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mymap[i][j], " ")
		}
		fmt.Println("")
	}
}
