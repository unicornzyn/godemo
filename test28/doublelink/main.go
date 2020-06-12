package main

import (
	"fmt"
)

// HeroNode struct
type HeroNode struct {
	no   int
	name string
	pre  *HeroNode
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
	node.pre = tmpNode
}

// link添加 有序
func insert2(head *HeroNode, node *HeroNode) {
	tmpNode := head
	for {
		if tmpNode.next == nil {
			break
		} else if tmpNode.next.no > node.no {
			break
		}

		tmpNode = tmpNode.next
	}

	node.next = tmpNode.next
	node.pre = tmpNode
	if tmpNode.next != nil {
		tmpNode.next.pre = node
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

// link显示 倒序
func list2(head *HeroNode) {
	if head.next == nil {
		fmt.Println("link为空")
		return
	}
	tmpNode := head
	for {
		if tmpNode.next == nil {
			break
		}
		tmpNode = tmpNode.next
	}

	for {
		if tmpNode.pre == nil {
			break
		}
		fmt.Printf("[%d, %s]=>", tmpNode.no, tmpNode.name)
		tmpNode = tmpNode.pre
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
		if tmpNode.next.next != nil {
			tmpNode.next.next.pre = tmpNode
		}
		tmpNode.next = tmpNode.next.next

	}
}

func main() {
	head := &HeroNode{}

	hero1 := &HeroNode{
		no:   1,
		name: "宋江",
	}

	hero2 := &HeroNode{
		no:   2,
		name: "卢俊义",
	}

	hero3 := &HeroNode{
		no:   3,
		name: "林冲",
	}

	insert2(head, hero1)
	insert2(head, hero3)
	insert2(head, hero2)
	list(head)
	list2(head)
	del(head, 2)
	list(head)
	list2(head)
}
