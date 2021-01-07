package Link

//头部操作  头部不存数据
type Node struct {
	Data  interface{}
	Pnext *Node
}

type LinkStack interface {
	IsEmpty() bool
	Push(data interface{})
	Pop() interface{}
	Length() int
}

func NewLinkStack() *Node {
	return &Node{}
}

func (ls *Node) IsEmpty() bool {
	return ls.Pnext == nil
}

//头部插入
func (ls *Node) Push(data interface{}) {
	newNode := &Node{Data: data}
	newNode.Pnext = ls.Pnext
	ls.Pnext = newNode

}

func (ls *Node) Pop() interface{} {
	if ls.IsEmpty() {
		return nil
	}
	value := ls.Pnext.Data
	ls.Pnext = ls.Pnext.Pnext
	return value
}

func (ls *Node) Length() int {
	pn := ls
	length := 0
	for pn.Pnext != nil {
		pn = pn.Pnext
		length++
	}
	return length

}
