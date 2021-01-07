package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	Remove()
	GetIndex() int
}

/*
 |         |  |
[0,1,2,3,4,5]
*/

type Iterable interface {
	Iterator() Iterator
}

type ArrayListIterator struct {
	list         *ArrayList
	currentIndex int
}

//实现接口

func (this *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator)
	it.list = this
	it.currentIndex = 0
	return it
}

func (this *ArrayListIterator) HasNext() bool {
	return this.currentIndex < this.list.TheSize

}

func (this *ArrayListIterator) Next() (interface{}, error) {
	if !this.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := this.list.Get(this.currentIndex)
	this.currentIndex++
	return value, err

}
func (this *ArrayListIterator) Remove() {
	this.currentIndex--
	this.list.Delete(this.currentIndex)

}
func (this *ArrayListIterator) GetIndex() int {
	return this.currentIndex
}
