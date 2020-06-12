package main

import (
	"errors"
	"fmt"
	"os"
)

// Queue 使用数组模拟队列
type Queue struct {
	maxSize int
	array   [5]int // 数组模拟队列
	head    int    //指向队列头部
	tail    int    //指向队列尾部
}

// Push 添加队列
func (q *Queue) Push(val int) (err error) {
	if q.IsFull() { //队列已满
		return errors.New("queue full")
	}
	q.array[q.tail] = val
	q.tail = (q.tail + 1) % q.maxSize
	return
}

// List 显示队列
func (q *Queue) List() {
	size := q.Size()
	if size == 0 {
		fmt.Println("队列为空")
		return
	}
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[(q.head+i)%q.maxSize])
	}
	fmt.Println("")
}

// Pop 取队列
func (q *Queue) Pop() (val int, err error) {
	if q.IsEmpty() {
		return -1, errors.New("queue empty")
	}
	val = q.array[q.head]
	q.head = (q.head + 1) % q.maxSize
	return
}

// IsFull 判断队列是否已满
func (q *Queue) IsFull() bool {
	return (q.tail+1)%q.maxSize == q.head
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.tail == q.head
}

// Size 判断队列有多少个元素
func (q *Queue) Size() int {
	return (q.tail + q.maxSize - q.head) % q.maxSize
}

func main() {
	q := &Queue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1. add 添加数据到队列")
		fmt.Println("2. get 从队列获取数据")
		fmt.Println("3. show 显示队列数据")
		fmt.Println("4. exit 退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要加入队列的数据:")
			fmt.Scanln(&val)
			err := q.Push(val)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("加入队列成功")
			}
		case "get":
			val, err := q.Pop()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("从队列中取出:", val)
			}
		case "show":
			q.List()
		case "exit":
			os.Exit(0)
		}
	}
}
