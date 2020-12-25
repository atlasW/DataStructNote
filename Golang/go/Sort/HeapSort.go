package main

import "fmt"

/* 堆排序   完全二叉树排序
5 > 3    &&  3 > 1  -->>5 > 1

*/

func HeapSortMax(arr []int) []int {
	lenth := len(arr)
	if lenth <= 1 {
		return arr
	}
	depth := lenth/2 - 1
	for i := depth; i >= 0; i-- {
		right := 2*i + 1
		left := 2*i + 2
		if right < lenth && arr[right] > arr[i] {
			arr[i], arr[right] = arr[right], arr[i]
		}
		if left < lenth && arr[left] > arr[i] {
			arr[i], arr[left] = arr[left], arr[i]
		}
		if right < lenth && left < lenth && arr[right] < arr[left] {
			arr[right], arr[left] = arr[left], arr[right]
		}

	}
	return arr
}

func main() {
	a := []int{23, 25, 734, 21, 1, 45, 2, 66, 3, 8, 9, 3, 4, 6}
	fmt.Println(HeapSortMax(a))
}
