package main

import "fmt"

/*
带head单链表例子,demo
*/

type HeroNode struct {
	No       int
	Name     string
	NickName string
	Next     *HeroNode
}

//显示链表所有信息
func ListSingle(head *HeroNode) {
	//先判断是否为空
	//辅助节点
	temp := head
	if temp.Next == nil {
		fmt.Println("empty single LinkedList")
	}
	for {
		fmt.Println(temp.Next)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}

}

//给链表增加一个节点
//在单链表最后加入
func InsertHeroNode(head *HeroNode, newheroNode *HeroNode) {
	//思路怎么找到最后节点
	//创建一个辅助节点跑龙套,定位链表最后位置
	temp := head
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	//新节点加入链表队尾
	temp.Next = newheroNode
}

//在链表合适位置插入 //经典案例
func InsertHeroNode2(head *HeroNode, newheroNode *HeroNode) {
	temp := head
	//思路需要清晰
	for {
		if temp.Next == nil { //若干为空链表,直接出去
			break
		} else if newheroNode.No >= temp.Next.No { //不满足条件 辅助节点右移.temp永远在不满足条件的前一位
			temp = temp.Next
		} else { //满足条件直接插入
			break
		}
	}
	//插入数据
	newheroNode.Next = temp.Next
	temp.Next = newheroNode
}

//给链表减少一个节点

func main() {
	//init singleNodeLinked
	//创建一个头节点
	head := &HeroNode{}
	//创建一个新节点
	hero1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}
	hero2 := &HeroNode{
		No:       2,
		Name:     "卢俊义",
		NickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		No:       3,
		Name:     "林冲",
		NickName: "豹子头",
	}
	hero4 := &HeroNode{
		No:       3,
		Name:     "吴用",
		NickName: "智多星",
	}

	//
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero4)
	InsertHeroNode2(head, hero2)
	ListSingle(head)

}
