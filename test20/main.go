package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//带缓冲区的文件读取
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') //读到\n结束
		if err == io.EOF {                  // io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}

	fmt.Println("文件读取结束")

	//一次性读取全部内容的方法
	content, err2 := ioutil.ReadFile("test.txt")
	if err2 != nil {
		fmt.Println("read file err=", err2)
	}
	fmt.Printf("%v", string(content)) //content 是 []byte 需要转为string输出
}
