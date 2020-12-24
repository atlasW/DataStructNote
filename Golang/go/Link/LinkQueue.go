package Link

//链式队列

type LinkQueue interface {
	Length() int
	DeQueue()
	EnQueue(data interface{})
}

type QuereLink struct {
	rear  *Node
	front *Node
}

func NewLinkQueue() *QuereLink {
	return &QuereLink{}
}

func (qlk *QuereLink) Length() int {
	pn := qlk.front
	length := 1
	if qlk.front == qlk.rear {
		return 0
	}
	for pn.Pnext != nil {
		pn = pn.Pnext
		length++
	}
	return length

}

func (qlk *QuereLink) EnQueue(data interface{}) {
	newnode := &Node{data, nil}
	if qlk.front == nil {
		qlk.front = newnode
		qlk.rear = newnode
	} else {
		qlk.rear.Pnext = newnode
		qlk.rear = qlk.rear.Pnext
	}

}

func (qlk *QuereLink) DeQueue() interface{} {
	resultNode := qlk.front
	if qlk.front == nil {
		return nil
	} else if qlk.front == qlk.rear {
		qlk.front = nil
		qlk.rear = nil
	} else {
		qlk.front = qlk.front.Pnext
	}
	return resultNode.Data

}
