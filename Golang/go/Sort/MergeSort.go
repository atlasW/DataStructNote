package main

import "fmt"

//归并操作（merge），也叫归并算法，指的是将两个已经排序的序列合并成一个序列的操作。归并排序算法依赖归并操作。   实质

// 递归   top -- down
// 迭代   bottom -- up

//迭代
func MergeSort(arr []int) []int {
	lenth := len(arr)
	if lenth < 2 {
		return arr
	}
	middle := len(arr) / 2
	left := MergeSort(arr[0:middle])
	right := MergeSort(arr[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	newArray := make([]int, len(left)+len(right))
	i, j, index := 0, 0, 0
	for {
		if left[i] < right[j] {
			newArray[index] = left[i]
			i++
			index++
		} else {
			newArray[index] = right[j]
			j++
			index++
		}
		//
		if i == len(left) {
			copy(newArray[index:], right[j:])
			break
		}
		if j == len(right) {
			copy(newArray[index:], left[i:])
			break
		}
	}
	return newArray

}

func main() {
	a := []int{25, 734, 21, 1, 45, 2, 66, 3, 8, 9, 3, 4, 6}
	fmt.Println(a)
	fmt.Println(MergeSort(a))
}
