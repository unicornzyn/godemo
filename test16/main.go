package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// Hero 定义一个结构体
type Hero struct {
	Name string
	Age  int
}

// HeroSlice 定义一个结构体切片
type HeroSlice []Hero

// 实现排序需要的三个方法 即实现了Sort的Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄～%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heros = append(heros, hero)
	}

	fmt.Println("排序前:")
	for _, v := range heros {
		fmt.Println(v)
	}
	sort.Sort(heros)
	fmt.Println("排序后:")
	for _, v := range heros {
		fmt.Println(v)
	}
}
