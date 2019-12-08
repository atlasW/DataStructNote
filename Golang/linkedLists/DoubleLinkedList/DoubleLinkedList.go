package main

import "fmt"

/*
双向链表例子,demo
梁山好汉排行榜
*/

type HeroNode struct {
	No       int
	Name     string
	NickName string
	Pre      *HeroNode
	Next     *HeroNode
}

//显示链表所有信息
func ListSingle(head *HeroNode) {
	//先判断是否为空
	//辅助节点
	temp := head
	if temp.Next == nil {
		fmt.Println("empty single LinkedList")
		return
	}
	for {
		fmt.Println(temp.Next)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}

}

//逆序输出
func ListSingle2(head *HeroNode) {
	//先判断是否为空
	//辅助节点
	temp := head
	if temp.Next == nil {
		fmt.Println("empty single LinkedList")
		return
	}
	//定位temp到最后节点
	for {
		if temp.Next == nil {
			break
		}
		temp = temp.Next
	}
	for {
		fmt.Println(temp)
		temp = temp.Pre
		if temp.Pre == nil {
			break
		}
	}

}

//给链表增加一个节点
//在链表最后加入
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
	newheroNode.Pre = temp
}

//在链表合适位置插入 //经典案例
func InsertHeroNode2(head *HeroNode, newheroNode *HeroNode) {
	temp := head
	//思路需要清晰
	flag := true
	for {
		if temp.Next == nil { //若干为空链表,直接出去
			break
		} else if newheroNode.No > temp.Next.No { //不满足条件 辅助节点右移.temp永远在不满足条件的前一位
			temp = temp.Next
		} else if newheroNode.No == temp.Next.No { //不允许同id插入
			flag = false
			break
		} else { //满足条件直接插入
			break
		}
	}
	//插入数据
	if !flag {
		fmt.Println("主键冲突")
		return
	}
	newheroNode.Next = temp.Next
	newheroNode.Pre = temp
	temp.Next = newheroNode
	if newheroNode.Next != nil {
		newheroNode.Next.Pre = newheroNode
	}
}

//给链表删除一个节点 //只适用没有重复id的链表
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	for {
		if temp.Next == nil { //找到最后一位
			break
		} else if id != temp.Next.No { //不满足条件 辅助节点右移.temp永远在不满足条件的前一位
			temp = temp.Next
		} else { //满足条件
			break
		}
	}
	//判断是否有需要删除的id
	if temp.Next != nil {
		temp.Next = temp.Next.Next
		if temp.Next != nil {
			temp.Next.Pre = temp
		}
	} else {
		fmt.Println("没有该需要删除的id")
	}
}

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
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero4)
	DelHeroNode(head, 2)
	fmt.Println()
	//  顺序输出
	ListSingle(head)
	//  逆序输出
	//ListSingle2(head)
}
