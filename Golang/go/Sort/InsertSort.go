package sort

// 1 9 2 8 4 6 7 5
// 1
// 1 9
// 1 2 9
// 1 2 8 9
//冒泡升级版

func InsertionSort(data []int) {
	b := len(data)
	if b == 1 {
		return
	}
	for i := 1; i < b; i++ {
		//最好情况下，数组已经是有序的，每插入一个元素，只需要考查前一个元素，因此最好情况下，插入排序的时间复杂度为O(N)。
		//   data[j] < data[j-1] 这一步就失败了
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

func InsertionSort2(data []int) {
	b := len(data)
	if b == 1 {
		return
	}
	for i := b - 2; i >= 0; i-- {
		for j := i; j <= b-2 && data[j] < data[j+1]; j++ {
			data[j], data[j+1] = data[j+1], data[j]

		}
	}
}

func InsertionSortString(data []string) {
	b := len(data)
	if b == 1 {
		return
	}
	for i := 1; i < b; i++ {
		//最好情况下，数组已经是有序的，每插入一个元素，只需要考查前一个元素，因此最好情况下，插入排序的时间复杂度为O(N)。
		//   data[j] < data[j-1] 这一步就失败了
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}
