package StackArray

type StackArray interface {
	Clean()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	dataSource  []interface{}
	capSize     int
	currentSize int
}

func Newstack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = make([]interface{}, 0, 1000)
	myStack.capSize = 1000
	myStack.currentSize = 0
	return myStack

}

func (st *Stack) Clean() {
	st.dataSource = make([]interface{}, 0, 1000)
	st.currentSize = 0
	st.capSize = 1000

}

func (st *Stack) Size() int {
	return st.currentSize
}

func (st *Stack) Pop() interface{} {
	if st.IsEmpty() {
		return nil
	}
	last := st.dataSource[st.currentSize-1]
	st.dataSource = st.dataSource[:st.currentSize-1]
	st.currentSize--
	return last

}

func (st *Stack) Push(data interface{}) {
	if !st.IsFull() {

		st.dataSource = append(st.dataSource, data)
		st.currentSize++

	}
}

func (st *Stack) IsFull() bool {
	if st.currentSize >= st.capSize {
		return true
	}
	return false

}

func (st *Stack) IsEmpty() bool {
	if st.currentSize == 0 {
		return true
	}
	return false
}
