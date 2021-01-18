package search

import sort "atlas/Sort"

/*
顺序搜索   遍历
*/

func BinSearch(arr []int, data int) int {
	sort.QuickSortPlus(arr)
	left := 0
	right := len(arr) - 1
	var mid int
	//二分查找
	for left <= right {
		mid = (left + right) / 2
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
