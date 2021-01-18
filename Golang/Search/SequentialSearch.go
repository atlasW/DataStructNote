package search

import sort "atlas/Sort"

/*
顺序搜索   遍历
*/

func SequentialSearch(arr []int, data int) int {
	sort.QuickSortPlus(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] == data {
			return i
		}
	}
	return -1
}
