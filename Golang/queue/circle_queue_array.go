package main

import (
	"errors"
	"fmt"
	"os"
)

/*
用array数组实现环形队列
1.head包含数据，但是初始为空状态下例外。tail不包含数据
2.maxsize 为5的数组，只能存四个元素,如果全部存满无法判断 是empty还是full .
3.打印队列的时候用了一个辅助变量。
*/
type CircleQueue struct {
	MaxSize int
	Array   []int
	Head    int
	Tail    int
}

//判断是否为满
func (this *CircleQueue) IsFull() bool {
	return this.Head == (this.Tail+1)%this.MaxSize
}

//判断是否为空
func (this *CircleQueue) IsEmpty() bool {
	return this.Head == this.Tail
}

//初始化一个对列
func NewQueue(size int) CircleQueue {
	this := CircleQueue{}
	array := make([]int, size)
	this.Array = array
	this.MaxSize = size
	this.Head = 0
	this.Tail = 0
	return this
}

//入队列
func (this *CircleQueue) InQueue(n int) error {
	if this.IsFull() {
		return errors.New("队列满了.")
	}
	this.Array[this.Tail] = n
	this.Tail = (this.Tail + 1) % this.MaxSize
	return nil
}

//出队列
func (this *CircleQueue) OutQueue() (value int, err error) {
	if this.IsEmpty() {
		err = errors.New("队列空了")
		return
	}
	value = this.Array[this.Head]
	this.Head = (this.Head + 1) % this.MaxSize
	return
}

//队列长度
func (this *CircleQueue) Size() int {
	return (this.Tail + this.MaxSize - this.Head) % this.MaxSize
}

// 打印出队列
func (this *CircleQueue) ShowInfo() {
	if this.IsEmpty() {
		fmt.Println("队列已空")
	}
	//借用辅助变量
	temp := this.Head
	for i := 0; i < this.Size(); i++ {
		fmt.Printf("arrar[%d]=%d\t", temp, this.Array[temp])
		temp = (temp + 1) % this.MaxSize
	}

}

func main() {
	queue := NewQueue(4)
	var key string
	var val int
	for {
		fmt.Println("1. 输入add 表示添加数据到队列")
		fmt.Println("2. 输入get 表示从队列获取数据")
		fmt.Println("3. 输入show 表示显示队列")
		fmt.Println("4. 输入exit 表示显示队列")

		fmt.Print("请输入:")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.InQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {

				fmt.Println("加入队列ok")
			}
		case "get":
			val, err := queue.OutQueue()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "show":
			queue.ShowInfo()
			fmt.Println()
		case "exit":
			os.Exit(0)
		}
	}

}
