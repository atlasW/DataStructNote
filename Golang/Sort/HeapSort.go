package sort

import "fmt"

/* 堆排序   完全二叉树排序
5 > 3    &&  3 > 1  -->>5 > 1

*/

// Heap get max  of  slice
func HeapSortMax(arr []int) {
	lenth := len(arr)
	if lenth <= 1 {
		return
	}
	// 树的深度
	depth := lenth/2 - 1
	for i := depth; i >= 0; i-- {
		right := 2*i + 1
		left := 2*i + 2
		//防止越界
		if right < lenth && arr[right] > arr[i] {
			arr[i], arr[right] = arr[right], arr[i]
		}
		if left < lenth && arr[left] > arr[i] {
			arr[i], arr[left] = arr[left], arr[i]
		}
		// if left < lenth && arr[right] < arr[left] {
		// 		arr[right], arr[left] = arr[left], arr[right]
		// 	}
	}
}

// HeapSort
func HeapSort(arr []int) {
	lenth := len(arr)
	//不能等于1
	for i := lenth - 1; i > 0; i-- {
		fmt.Printf("%s --max:%s--%+v\n", i, arr[i], arr)
		// 这里特别重要    包括前面 不包括后面...... 这里耽误了很长时间。。。。
		HeapSortMax(arr[0 : i+1])
		fmt.Printf("%s --max:%s--%+v\n", i, arr[i], arr)
		arr[0], arr[i] = arr[i], arr[0]
		fmt.Printf("%s --max:%s--%+v\n", i, arr[i], arr)
		fmt.Println()
	}

}

//
//func main() {
//	a := []int{23, 25, 734, 21, 1, 45, 2, 66, 3, 8, 9, 3, 4, 6}
//	fmt.Println(a)
//	HeapSort(a)
//	fmt.Println(a)
//}
