package main

import (
	"atlas/ArrayList"
	"fmt"
)

func main() {
	var list ArrayList.List = ArrayList.NewArrayList()
	list.Append("a1")
	list.Append("b2")
	list.Append("c3")
	list.Append("c3")
	list.Append("c3")
	for it := list.Iterator(); it.HasNext(); {
		item, _ := it.Next()
		fmt.Println(item)
	}
	fmt.Println(list)
}
