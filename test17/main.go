package main

import "fmt"

// Usb 接口
type Usb interface {
	Start()
	Stop()
}

// Camera 相机结构体
type Camera struct {
	name string
}

// Start 实现接口方法
func (c Camera) Start() {
	fmt.Println("相机", c.name, "开始工作")
}

// Stop 实现接口方法
func (c Camera) Stop() {
	fmt.Println("相机", c.name, "停止工作")
}

// Phone 手机结构体
type Phone struct {
	name string
}

// Start 实现接口方法
func (p Phone) Start() {
	fmt.Println("手机", p.name, "开始工作")
}

// Stop 实现接口方法
func (p Phone) Stop() {
	fmt.Println("手机", p.name, "停止工作")
}

// Call Phone结构体独有方法
func (p Phone) Call() {
	fmt.Println("手机", p.name, "开始呼叫")
}

// Computer 计算机结构体
type Computer struct {
}

// Working 计算机工作方法 传入Usb接口类型
func (c Computer) Working(usb Usb) {
	usb.Start()
	//使用类型断言 调用Phone独有的Call方法
	if p, ok := usb.(Phone); ok {
		p.Call()
	}
	usb.Stop()
}

// TypeJudge 判断传入参数的类型 items可变长度参数 interface{} 可接收任意类型
func TypeJudge(items ...interface{}) {
	for i, x := range items {
		switch x.(type) { //这里的type是关键字 固定写法
		case bool:
			fmt.Printf("params #%d is a bool 值是%v\n", i, x)
		case float64:
			fmt.Printf("params #%d is a float64 值是%v\n", i, x)
		case int, int64:
			fmt.Printf("params #%d is a int 值是%v\n", i, x)
		case nil:
			fmt.Printf("params #%d is a nil 值是%v\n", i, x)
		case string:
			fmt.Printf("params #%d is a string 值是%v\n", i, x)
		case Computer:
			fmt.Printf("params #%d is a Computer 值是%v\n", i, x)
		default:
			fmt.Printf("params #%d is a unknown 值是%v\n", i, x)
		}
	}
}

func main() {
	// 类型断言示例1
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	computer := Computer{}
	for _, v := range usbArr {
		computer.Working(v)
	}

	// 类型断言示例2
	TypeJudge(1, 2.1, "abc", nil, true, computer)
}
