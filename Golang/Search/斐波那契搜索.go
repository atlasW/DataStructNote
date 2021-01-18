package search

import (
	sort "atlas/Sort"
	"fmt"
)

/*
  中值搜索 //差距比较小的
*/

func BinSearchMid(arr []int, data int) int {
	sort.QuickSortPlus(arr)
	left := 0
	right := len(arr) - 1
	var mid int
	//二分查找
	for left <= right {
		//mid = (left + right) / 2  中值查找
		leftv := float64(data - arr[left])
		allv := float64(arr[right] - arr[left])
		diff := float64(right - left)
		mid = int(float64(left) + (leftv/allv)*diff)
		fmt.Println(mid)
		if mid > right || mid < left {
			return -1
		}
		if arr[mid] > data {
			right = mid - 1
		} else if arr[mid] < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
