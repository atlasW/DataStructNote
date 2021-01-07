package sort

import "fmt"

//冒泡排序   鸡尾酒排序

func CocktailBubbleSort(arr []int) {
	left := 0
	right := len(arr) - 1
	swapped := true
	if left < right && swapped {
		swapped = false
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		right--
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		left++
		fmt.Println(arr)
	}
}

//经典冒泡
func BubbleSort(arr []int) {
	swapped := true
	top := len(arr) - 1
	for swapped {
		swapped = false
		for i := 0; i < top; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		top--
	}
}
