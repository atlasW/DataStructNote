package ArrayList

type StackXArray interface {
	Clean()
	Size() int
	Pop() interface{}
	Push(data interface{})
	IsFull() bool
	IsEmpty() bool
}

type StackX struct {
	dataSource *ArrayList
	Myit       Iterator
}

func NewArrayListStackX() *StackX {
	myStackX := new(StackX)
	myStackX.dataSource = NewArrayList()
	myStackX.Myit = myStackX.dataSource.Iterator()
	return myStackX

}

func (st *StackX) Clean() {
	// st.dataSource = make([]interface{}, 0, 10)
	st.dataSource.Clean()
	st.dataSource.TheSize = 0

}

func (st *StackX) Size() int {
	return st.dataSource.TheSize
}

func (st *StackX) Pop() interface{} {
	if st.IsEmpty() {
		return nil
	}
	last := st.dataSource.dataStore[st.dataSource.TheSize-1]
	st.dataSource.Delete(st.dataSource.TheSize - 1)
	return last

}

func (st *StackX) Push(data interface{}) {
	if !st.IsFull() {
		st.dataSource.Append(data)

	}
}

func (st *StackX) IsFull() bool {
	if st.dataSource.TheSize >= 10 {
		return true
	}
	return false

}

func (st *StackX) IsEmpty() bool {
	if st.dataSource.TheSize == 0 {
		return true
	}
	return false
}
