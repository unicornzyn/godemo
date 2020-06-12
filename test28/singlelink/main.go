package main

import (
	"fmt"
)

// HeroNode struct
type HeroNode struct {
	no   int
	name string
	next *HeroNode
}

// link添加 尾部追加
func insert(head *HeroNode, node *HeroNode) {
	tmpNode := head
	for {
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	tmpNode.next = node
}

// link添加 有序
func insert2(head *HeroNode, node *HeroNode) {
	tmpNode := head
	for {
		if tmpNode.next == nil {
			break
		} else if tmpNode.next.no > node.no {
			node.next = tmpNode.next
			break
		}
		tmpNode = tmpNode.next
	}

	tmpNode.next = node
}

// link显示
func list(head *HeroNode) {
	if head.next == nil {
		fmt.Println("link为空")
		return
	}
	tmpNode := head

	for {
		fmt.Printf("[%d, %s]=>", tmpNode.next.no, tmpNode.next.name)
		tmpNode = tmpNode.next
		if tmpNode.next == nil {
			break
		}

	}
	fmt.Println("")
}

// 删除link
func del(head *HeroNode, no int) {
	tmpNode := head
	flag := false
	for {
		if tmpNode.next == nil {
			break
		} else if tmpNode.next.no == no {
			flag = true
			break
		}
		tmpNode = tmpNode.next
	}
	if flag {
		tmpNode.next = tmpNode.next.next
	}
}

func main() {
	head := &HeroNode{}

	hero3 := &HeroNode{
		no:   3,
		name: "林冲",
	}

	insert2(head, hero3)

	hero1 := &HeroNode{
		no:   1,
		name: "宋江",
	}

	insert2(head, hero1)

	hero2 := &HeroNode{
		no:   2,
		name: "卢俊义",
	}

	insert2(head, hero2)
	list(head)
	del(head, 2)
	list(head)
}
