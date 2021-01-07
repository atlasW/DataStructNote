package sort

// 1 9 2 8 4 6 7 5
// 9
// 9 8
// 9 8 7
//O(n^2)

func SelectSort(this []int) {
	lenth := len(this)
	if lenth <= 1 {
	} else {
		//只剩一个元素 不需要挑选
		for i := 0; i < lenth-1; i++ {
			min := i
			for j := i + 1; j < lenth; j++ {
				if this[min] > this[j] {
					min = j
				}
			}
			if i != min {
				this[i], this[min] = this[min], this[i]
			}
		}
	}
}

func SelectSortString(this []string) {
	lenth := len(this)
	if lenth <= 1 {
	} else {
		//只剩一个元素 不需要挑选
		for i := 0; i < lenth-1; i++ {
			min := i
			for j := i + 1; j < lenth; j++ {
				if this[min] > this[j] {
					min = j
				}
			}
			if i != min {
				this[i], this[min] = this[min], this[i]
			}
		}
	}
}
