package main

import (
	"fmt"
	"reflect"
)

// Monster 结构体声明
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

// Print 结构体方法
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

// GetSum 结构体方法
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// Set 结构体方法
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// TestStruct 反射的应用案例
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)

	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, val.Field(i))
		//获取struct的tag标签
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	// 调用的是第2个方法 方法不是通过代码实现顺序排序 而是根据方法的名字升序排序的
	val.Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是 []reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果是 []reflect.Value

}

func main() {
	a := Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}

	TestStruct(a)
}
