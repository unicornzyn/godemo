package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 字符串常用的函数
	str := "test12345中国"
	// go的编码统一 utf8   ascii字符占1个字节  汉字占3个字节 len按字节返回
	fmt.Println("字符串长度:", len(str))

	// 字符串遍历 解决中文问题
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c\n", r[i])
	}

	//字符串/整数互转
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转化错误", err)
	} else {
		fmt.Println("转化的结果是 ", n)
	}

	str2 := strconv.Itoa(111)
	fmt.Printf("str = %v, Type= %T\n", str2, str2)

	//字符串转byte[]
	var bytes = []byte("hello world, 中国")
	fmt.Printf("bytes = %v\n", bytes)

	//byte[]转字符串
	fmt.Println("str = ", string([]byte{97, 98, 99}))

	//10进制转 2 8 16进制 返回字符串
	var i int64 = 125
	fmt.Printf("10 => %v, 2 => %v, 8 => %v, 16 => %v\n", i, strconv.FormatInt(i, 2), strconv.FormatInt(i, 8), strconv.FormatInt(i, 16))

	// 判断子串包含
	fmt.Printf("seafood是否包含foo b=%v\n", strings.Contains("seafood", "foo"))

	//统计字符串中有几个子串
	fmt.Printf("abcdsacdf包含%v个cd\n", strings.Count("abcdsacdf", "cd"))

	//不区分大小写的比较  ==比较区分大小写
	fmt.Println("abc是否等于Abc b=", strings.EqualFold("abc", "Abc"))

	//返回子串在字符串中的第一次出现的index 未出现-1，最后一次出现的index
	fmt.Println("abcabca中ca的index=", strings.Index("abcabca", "ca"))
	fmt.Println("abcabca中ca的LastIndex=", strings.LastIndex("abcabca", "ca"))

	//字符串替换 最后一个参数 是替换的个数  -1为全部替换
	str3 := "Hello,world,world hello~"
	fmt.Println(str3, "world替换为中国=>", strings.Replace(str3, "world", "中国", 1))

	//字符串分割
	strArr := strings.Split(str3, ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}

	//字符串大小写转化
	fmt.Printf("%v => toupper(%v) => tolower(%v)\n", str3, strings.ToLower(str3), strings.ToUpper(str3))

	//去除左右两边的空格
	fmt.Printf("%q\n", strings.TrimSpace(" aa bb    "))

	//去除两边指定的字符
	fmt.Printf("%q\n", strings.Trim("!! hel!lo! ", " !"))

	//判断是否以指定的字符串开头和结尾
	fmt.Println("b1=", strings.HasPrefix("abcdefg", "abc"), "b2=", strings.HasSuffix("abcdefg", "fg"))
}
