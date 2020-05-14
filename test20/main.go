package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readfile() {
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

func writefile() {
	//带缓冲区的文件读取
	file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 3; i++ {
		writer.WriteString("hello world~~~\n")
	}

	//调用此方法才会把缓冲区的内容真正写入文件
	writer.Flush()
}

// PathExists 判断文件或路径是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	//readfile()
	//writefile()
	if isExists, _ := PathExists("test.txt"); isExists {
		fmt.Println("test.txt 存在")
	} else {
		fmt.Println("test.txt 不存在")
	}
}
