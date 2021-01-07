package Queue

type MyQueue interface {
	Size() int
	Front() interface{}
	End() interface{}
	isEmpty() bool
	EnterQueue(data interface{})
	OutQueue(data interface{})
	Clear()
}

type Queue struct {
	dataStore []interface{}
	theSize   int
}

func NewQueue() *Queue {
	myqueue := new(Queue)
	myqueue.dataStore = make([]interface{}, 0)
	myqueue.theSize = 0
	return myqueue
}

func (myq *Queue) Size() int {
	return myq.theSize
}
func (myq *Queue) Front() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[0]

}
func (myq *Queue) End() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	return myq.dataStore[myq.Size()-1]

}
func (myq *Queue) isEmpty() bool {
	return myq.Size() == 0
}
func (myq *Queue) EnterQueue(data interface{}) {
	myq.dataStore = append(myq.dataStore, data)
	myq.theSize++
}
func (myq *Queue) OutQueue() interface{} {
	if myq.Size() == 0 {
		return nil
	}
	data := myq.dataStore[0]
	if myq.Size() > 1 {
		myq.dataStore = myq.dataStore[1:myq.Size()]
		myq.theSize--
	} else {
		myq.dataStore = make([]interface{}, 0)
		myq.theSize--
	}
	return data

}
func (myq *Queue) Clear() {
	myq.dataStore = make([]interface{}, 0)
	myq.theSize = 0
}
