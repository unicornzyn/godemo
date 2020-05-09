package main

import "fmt"

func main() {
	// map使用方式1
	var citys1 map[string]string
	citys1 = make(map[string]string, 3)
	citys1["no1"] = "北京"
	citys1["no2"] = "上海"
	citys1["no3"] = "天津"
	fmt.Println("citys1: ", citys1)

	// map使用方式2
	citys2 := make(map[string]string)
	citys2["no1"] = "北京"
	citys2["no2"] = "上海"
	citys2["no3"] = "天津"
	fmt.Println("citys2: ", citys2)

	// map使用方式3
	citys3 := map[string]string{
		"no1": "北京",
		"no2": "上海",
		"no3": "天津",
	}
	fmt.Println("citys3: ", citys3)

	// 添加一个新的值
	citys3["no4"] = "重庆"
	fmt.Println("添加后的city3: ", citys3)

	// 修改值
	citys3["no2"] = "上海～"
	fmt.Println("修改后的citys3: ", citys3)

	// 删除 删除一个不存在的key不会报错 想全部删除key 需要遍历依次删除 或者使用make重新给一个新的 原值等待gc回收
	delete(citys3, "no3")
	fmt.Println("删除后的citys3: ", citys3)

	// 查找
	val, hasVal := citys3["no3"]
	if hasVal {
		fmt.Println("存在val=", val)
	} else {
		fmt.Println("不存在")
	}

	// 遍历
	for k, v := range citys3 {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
}
