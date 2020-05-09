package main

import (
	"errors"
	"fmt"
)

//异常处理 defer panic recover

func test() {
	defer func() {
		//recover是内置函数 可以捕获异常
		if err := recover(); err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

// test2 自定义错误
func test2(name string) (err error) {
	if name == "a" {
		return nil
	} else {
		return errors.New("自定义错误")
	}
}

func main() {
	test()
	fmt.Println("over~")
	err := test2("b")
	if err != nil {
		//输出错误并终止程序
		panic(err)
	}
	fmt.Println("ok~")
}
