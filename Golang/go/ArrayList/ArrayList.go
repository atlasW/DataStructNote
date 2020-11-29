package ArrayList

import (
	"errors"
	"fmt"
)

//接口
type List interface {
	Size() int
	Get(index int) (interface{}, error)
	Set(index int, newval interface{}) error
	Insert(index int, newval interface{}) error
	Append(newval interface{})
	Clean()
	Delete(index int) error
	String() string
	Iterator() Iterator
}

//数组结构
type ArrayList struct {
	dataStore []interface{}
	TheSize   int
}

func NewArrayList() *ArrayList {
	list := new(ArrayList)
	list.dataStore = make([]interface{}, 0, 10)
	list.TheSize = 0
	return list
}

func (this *ArrayList) Size() int {
	return this.TheSize
}

func (this *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= this.TheSize {
		return nil, errors.New("索引越界")

	}
	return this.dataStore[index], nil
}

func (this *ArrayList) Append(newval interface{}) {
	this.dataStore = append(this.dataStore, newval)
	this.TheSize++
}

func (this *ArrayList) String() string {
	return fmt.Sprint(this.dataStore)
}

func (this *ArrayList) Set(index int, newval interface{}) error {
	if index < 0 || index >= this.TheSize {
		return errors.New("索引越界")

	}
	this.dataStore[index] = newval
	return nil
}

func (this *ArrayList) Clean() {
	this.dataStore = make([]interface{}, 0, 10)
	this.TheSize = 0
}

//重新开辟内存
func (this *ArrayList) Insert(index int, newval interface{}) error {
	if index < 0 || index >= this.TheSize {
		return errors.New("索引越界")

	}
	this.checksiFull()                               //检测内存 自动扩容
	this.dataStore = this.dataStore[:this.TheSize+1] //
	for i := this.TheSize; i > index; i-- {
		this.dataStore[i] = this.dataStore[i-1]
	}
	this.dataStore[index] = newval
	this.TheSize++
	return nil
}

func (this *ArrayList) Delete(index int) error {
	if index < 0 || index >= this.TheSize {
		return errors.New("索引越界")

	}
	this.dataStore = append(this.dataStore[0:index], this.dataStore[index+1:])
	this.TheSize--
	return nil
}

func (this *ArrayList) checksiFull() {
	if this.TheSize == cap(this.dataStore) {
		//注意这地方
		newdateStore := make([]interface{}, this.TheSize, 2*this.TheSize)
		copy(newdateStore, this.dataStore)
		// 	for i := 0; i < len(this.dataStore); i++ {
		// 		newdateStore[i] = this.dataStore[i]
		// 	}
		this.dataStore = newdateStore
	}
}
