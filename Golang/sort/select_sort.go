package main

import (
	"fmt"
	"time"
)

func main() {
	startT := time.Now() //计算当前时间

	//	a := []int{5, 2, 3, 56, 21, 111, 333, 55, 12, 1, 22}
	a := []int{1, 2, 3, 4, 5, 7, 6}
	//SelectSort(a)
	//	selection_sort(a)
	fmt.Println(a)
	Bubble_Sort3(a)
	fmt.Println(a)

	tc := time.Since(startT) //计算耗时
	fmt.Printf("time cost = %v\n", tc)
}

// slice 是指针传递  array 是值传递
//select sort     少了一个临时变量记录max 合不合适
// O(n^2)
func SelectSort(array []int) {
	//第一层for  到了 倒数第二个数截止
	for i := 0; i < len(array)-1; i++ {
		//第二层for  从i+1到最后和array[i]对比
		for i2 := i + 1; i2 <= len(array)-1; i2++ {
			if array[i2] > array[i] {
				array[i2], array[i] = array[i], array[i2]
			}

		}

	}
}

//冒泡 ? ?
func Bubble_Sort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//误区了 。。。 最大或者最小的数 堆积在了右边  而不在左边。。。
		// 	for i2 := i; i2 < len(array)-1; i2++ {
		// 		if array[i2] < array[i2+1] {
		// 			array[i2], array[i2+1] = array[i2+1], array[i2]
		// 		}
		for i2 := len(array) - 1; i2 > i; i2-- {
			if array[i2] < array[i2-1] {
				array[i2], array[i2-1] = array[i2-1], array[i2]
			}
		}

		fmt.Println(array)
	}
}

func Bubble_Sort2(array []int) {
	for i := len(array) - 1; i > 0; i-- {
		for i2 := 0; i2 < i; i2++ {
			if array[i2] > array[i2+1] {
				array[i2], array[i2+1] = array[i2+1], array[i2]
			}

		}
	}
}

func Bubble_Sort3(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		swap := 0
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap++
			}
		}
		fmt.Println(swap)
		if swap == 0 {
			break
		}
	}
	return arr
}

func Bubble_Sort4(arr []int) {
	swapped := true
	len := len(arr)
	for swapped {
		swapped = false
		for i := 0; i < len-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
}
