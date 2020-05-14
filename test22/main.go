package main

import (
	"encoding/json"
	"fmt"
)

//A 结构体
type A struct {
	B string `json:"b_key"`
	C int
	D bool
}

func main() {
	a := A{"bval", 100, true}
	fmt.Printf("原值=%v\n", a)
	v, err := json.Marshal(a)
	if err != nil {
		fmt.Println("序列化错误 err=", err)
		return
	}
	fmt.Printf("序列化值=%v\n", string(v))
	v2 := "{\"b_key\":\"bval2\",\"C\":1000,\"D\":false}"
	fmt.Printf("原值=%v\n", v2)
	var a2 A
	err2 := json.Unmarshal([]byte(v2), &a2)
	if err2 != nil {
		fmt.Println("反序列化错误 err=", err2)
		return
	}
	fmt.Printf("反序列化值=%v\n", a2)
}
