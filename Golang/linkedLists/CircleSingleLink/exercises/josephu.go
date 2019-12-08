package main

import (
	"fmt"
	"strconv"
)

//环形列表
type CatNode struct {
	no   int    // no
	name string //name
	next *CatNode
}

//insert
func InsertCat(head *CatNode, newcat *CatNode) {
	//first node
	if head.next == nil {
		//	head = newcat     // 不能这么写  改变了地址未知
		head.no = newcat.no
		head.name = newcat.name
		head.next = head
		return
	}

	// not first
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入链表中
	temp.next = newcat
	newcat.next = head
}

//show
func ShowCat(head *CatNode) {
	if head.next == nil {
		fmt.Println("空链表")
		return
	}

	temp := head
	for {
		fmt.Println(temp)
		if temp.next == head {
			break
		}
		temp = temp.next

	}

}

// delete  与案例不同
func DelCatNode(head *CatNode, id int) {
	temp := head
	//empty
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}

	//only one node
	if temp.next == head {
		if head.no == id {
			head.next = nil
		}
		return
	}
	// mult
	if head.no != id {
		for {
			if temp.next == head {
				fmt.Println("链表中没有需要删除的id")
				return
			}
			if temp.next.no == id {
				temp.next = temp.next.next
				break
			}
			temp = temp.next
		}

	} else {
		head.no = temp.next.no
		head.name = temp.next.name
		head.next = temp.next.next

	}

}

//约瑟夫问题
func ysfwt(head *CatNode, n int) {
	temp := head
	//empty
	if temp.next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		for i := 0; i < n-1; i++ {
			temp = temp.next
		}
		fmt.Println(temp.next.no)
		temp.next = temp.next.next
		if temp.next == temp {
			fmt.Println(temp.no)
			break
		}

	}
}

func main() {
	head := &CatNode{}

	for i := 1; i <= 100; i++ {
		InsertCat(head, &CatNode{
			no:   i,
			name: "Cat" + strconv.Itoa(i),
		})

	}
	ysfwt(head, 2)

}
