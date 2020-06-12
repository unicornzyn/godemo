package main

import (
	"fmt"
)

// Boy struct
type Boy struct {
	no   int
	next *Boy
}

// add 单向环形链表 num 个数  返回第一个boy的指针
func add(num int) *Boy {
	first := &Boy{}
	cur := &Boy{}

	for i := 1; i <= num; i++ {
		boy := &Boy{
			no: i,
		}
		if i == 1 {
			first = boy
			cur = boy
		}
		boy.next = first
		cur.next = boy
		cur = boy
	}
	return first
}

// show 显示环形链表
func show(first *Boy) {
	if first.next == nil {
		fmt.Println("空链表")
		return
	}

	curr := first
	for {
		fmt.Printf("%d->", curr.no)
		if curr.next == first {
			break
		}
		curr = curr.next
	}
	fmt.Println("")
}

// play 设编号为1,2,...,n的n个人围成一圈, 约定编号为k(1<=k<=n)
// 的人从1开始报数, 数到m的那个人出列, 他的下一位又从1开始报数,
// 数到m的那个人又出列, 依次类推, 直到所有人出列为止, 由此产生一个
// 出队编号的序列
func play(n int, k int, m int) {
	if n < 1 {
		fmt.Println("n要大于0的整数")
		return
	}

	// 生成环形链表
	first := add(n)

	// 显示链表
	show(first)

	tail := first
	// tail移动到first后面
	for {
		if tail.next == first {
			break
		}
		tail = tail.next
	}

	// first移动到k位置 tail跟随first移动继续保持在first后面
	for i := 0; i < k-1; i++ {
		first = first.next
		tail = tail.next
	}

	//依次出列
	for {
		// 移动m下 删除first
		for i := 0; i < m-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("%d->", first.no)
		first = first.next
		tail.next = first
		if tail == first {
			break
		}
	}
	// 最后一个出列
	fmt.Println(first.no)
}

// 约瑟夫问题
func main() {
	play(5, 2, 3)
}
