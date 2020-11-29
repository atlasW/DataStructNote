package ArrayList

type StackArray interface {
	Clean()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type Stack struct {
	dataSource *ArrayList
	capSize    int
}

func NewArrayListStack() *Stack {
	myStack := new(Stack)
	myStack.dataSource = NewArrayList()
	myStack.capSize = 10
	return myStack

}

func (st *Stack) Clean() {
	// st.dataSource = make([]interface{}, 0, 10)
	st.dataSource.Clean()
	st.capSize = 10

}

func (st *Stack) Size() int {
	return st.dataSource.TheSize
}

func (st *Stack) Pop() interface{} {
	if st.IsEmpty() {
		return nil
	}
	last := st.dataSource.dataStore[st.dataSource.TheSize-1]
	st.dataSource.Delete(st.dataSource.TheSize - 1)
	return last

}

func (st *Stack) Push(data interface{}) {
	if !st.IsFull() {
		st.dataSource.Append(data)

	}
}

func (st *Stack) IsFull() bool {
	if st.dataSource.TheSize >= st.capSize {
		return true
	}
	return false

}

func (st *Stack) IsEmpty() bool {
	if st.dataSource.TheSize == 0 {
		return true
	}
	return false
}
