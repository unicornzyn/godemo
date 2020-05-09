package main

import (
	"encoding/json"
	"fmt"
)

// Monster 自定义struct 同时测试strcut tag的使用
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"weapon"`
}

// Say 方法
func (m Monster) Say() {
	fmt.Println("我的名字是", m.Name)
}

// String 发放会被fmt.Println自动调用
// func (m Monster) String() string {
func (m *Monster) String() string {
	return "Name:" + m.Name
}

func main() {
	monster := Monster{"牛魔王", 500, "芭蕉扇～"}

	// 测试struct tag
	jsonstr, _ := json.Marshal(monster)
	fmt.Println("jsonstr", string(jsonstr))

	// 调用方法
	monster.Say()

	// 自动调用String() 如果定义的是 *Monster指针类型 则调用时 使用&取地址符才会自动调用String()
	// fmt.Println(monster)
	fmt.Println(&monster)
}
