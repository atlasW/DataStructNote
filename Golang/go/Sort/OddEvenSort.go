package sort

func OddEvenSort(arr []int) {
	length := len(arr)
	needSorted := true
	for needSorted == true {
		needSorted = false
		for i := 1; i < length-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				needSorted = true
			}

		}
		for i := 0; i < length-1; i = i + 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				needSorted = true
			}
		}

	}

}
