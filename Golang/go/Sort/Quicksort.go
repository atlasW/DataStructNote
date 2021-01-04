package sort

/*
 function partition(a, left, right, pivotIndex)
 {
     pivotValue = a[pivotIndex]
     swap(a[pivotIndex], a[right]) // 把pivot移到結尾
     storeIndex = left
     for i from left to right-1
     {
         if a[i] <= pivotValue
          {
             swap(a[storeIndex], a[i])
             storeIndex = storeIndex + 1
          }
     }
     swap(a[right], a[storeIndex]) // 把pivot移到它最後的地方
     return storeIndex
 }
 procedure quicksort(a, left, right)
     if right > left
         select a pivot value a[pivotIndex]
         pivotNewIndex := partition(a, left, right, pivotIndex)
         quicksort(a, left, pivotNewIndex-1)
         quicksort(a, pivotNewIndex+1, right)

*/
//分割
func partition(a []int, left, right, pivotIndex int) int {
	//基准值
	pivotValue := a[pivotIndex]
	// 基准和最后一位 互换
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	storeIndex := left
	//找到基准所在的位置   //很巧妙的方法  找到基准的位置  并且将 小于基准的移动到基准的左边
	for i := left; i < right; i++ {
		if a[i] <= pivotValue {
			a[storeIndex], a[i] = a[i], a[storeIndex]
			storeIndex++
		}

	}
	a[right], a[storeIndex] = a[storeIndex], a[right]
	return storeIndex

}
func QuickSort(a []int, left, right int) {
	if right > left {
		pivotIndex := left
		pivotNewIndex := partition(a, left, right, pivotIndex)
		QuickSort(a, left, pivotNewIndex-1)
		QuickSort(a, pivotNewIndex+1, right)
	}
}
