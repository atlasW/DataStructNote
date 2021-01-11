package main

import (
	search "atlas/Search"
	sort "atlas/Sort"
	"fmt"
)

//func main() {
//	var list ArrayList.List = ArrayList.NewArrayList()
//	list.Append("a1")
//	list.Append("b2")
//	list.Append("c3")
//	list.Append("c3")
//	list.Append("c3")
//	for it := list.Iterator(); it.HasNext(); {
//		item, _ := it.Next()
//		fmt.Println(item)
//	}
//	fmt.Println(list)
//}

//func main() {
//	myst := ArrayList.NewArrayListStackX()
//	myst.Push(11)
//	myst.Push(22)
//	myst.Push(33)
//	myst.Push(44)
//	//iterator
//	for it := myst.Myit; it.HasNext(); {
//		item, _ := it.Next()
//		fmt.Println(item)
//	}
//	fmt.Println(myst.Pop())
//}
// 1 1
// 2 1
// 3 2
// 4 3
// 5 5
// 6 8
//斐波那契数列  栈实现
// func main() {
// 	mystack := StackArray.Newstack()
// 	mystack.Push(7)
// 	last := 0
// 	for !mystack.IsEmpty() {
// 		data := mystack.Pop()
// 		if data == 1 || data == 2 {
// 			last += 1
// 		} else {
// 			mystack.Push(data.(int) - 1)
// 			mystack.Push(data.(int) - 2)
// 		}
// 	}
// 	fmt.Println(last)
// }

//递归文件夹  堆栈实现stack
// func main() {
// 	mystack := StackArray.Newstack()
// 	path := "/Users/atlas/atlas/go/develop/DataStructNote/Golang"
// 	mystack.Push(path)
// 	pathArray := []string{}
// 	for !mystack.IsEmpty() {
// 		getpath := mystack.Pop()
// 		pathArray = append(pathArray, getpath.(string))
// 		files, _ := ioutil.ReadDir(getpath.(string))
// 		for _, file := range files {
// 			fullpath := getpath.(string) + "/" + file.Name()
// 			if file.IsDir() {
// 				mystack.Push(fullpath)
//
// 			} else {
// 				pathArray = append(pathArray, fullpath)
//
// 			}
//
// 		}
//
// 	}
// 	for _, v := range pathArray {
// 		fmt.Println(v)
// 	}
// }

//递归文件夹  队列实现stack
// 堆栈访问文件夹  深度遍历  队列实现广度遍历

//循环队列
//func main() {
//	var myq CricleQueue.CricleQueue
//	CricleQueue.InitQueue(&myq)
//	for i := 1; i <= 99; i++ {
//		err := CricleQueue.EnQueue(&myq, i)
//		if err != nil {
//			panic(err)
//		}
//	}
//}

// func main() {
// 	linkstack := Link.NewLinkStack()
// 	for i := 0; i < 100000000; i++ {
// 		linkstack.Push(i)
// 	}
//
// 	for data := linkstack.Pop(); data != nil; data = linkstack.Pop() {
// 		fmt.Println(data)
// 	}
// }
//
//func main() {
//	linkqueue := Link.NewLinkQueue()
//	for i := 0; i < 100; i++ {
//		linkqueue.EnQueue(i)
//	}
//
//	// for i := 0; i < 99; i++ {
//	// 	fmt.Println(linkqueue.DeQueue())
//	// }
//	fmt.Println(linkqueue.Length())
//
//}

//selectsort
func main() {
	a := []int{2, 4, 5, 3, 1, 2, 4, 5, 7, 8, 4}
	//Sort.SelectSort(a)
	fmt.Println(a)
	//Sort.SelectSortString(b)
	sort.ShellSort(a)
	fmt.Println(search.BinSearch(a, 5))

}
