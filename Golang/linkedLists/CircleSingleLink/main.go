package main

import "fmt"

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

func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	InsertCat(head, cat1)
	InsertCat(head, cat2)
	InsertCat(head, cat3)
	ShowCat(head)
	DelCatNode(head, 2)
	DelCatNode(head, 30)
	fmt.Println("删除后")
	ShowCat(head)
}
