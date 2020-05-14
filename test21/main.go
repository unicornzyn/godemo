package main

import (
	"flag"
	"fmt"
)

func main() {
	var a string
	var bb int
	var c string
	//参数说明 &a 接收参数 "a" 命令行参数格式 "" 默认值 "说明" 参数说明
	flag.StringVar(&a, "a", "", "说明")
	flag.IntVar(&bb, "bb", 0, "bb")
	flag.StringVar(&c, "c", "default c", "c")

	//必须调用这个方法(os.Args里是原始参数切片)
	flag.Parse()

	// go run main.go -a test -bb 20
	fmt.Printf("a=%v, bb=%v, c=%v\n", a, bb, c)

}
