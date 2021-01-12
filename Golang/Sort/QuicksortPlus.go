package sort

//改良版本的快速排序  解约内存
func QuickSortPlus(arr []int) {
	QuickSortX(arr, 0, len(arr)-1)
}

//插入排序
func SortForMerge(arr []int, left, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i] //备份数据
		var j int
		for j = i; j > left && arr[j-1] > temp; j-- { //定位
			arr[j] = arr[j-1] //数据后移
		}
		arr[j] = temp
	}

}

//递归快速排序
func QuickSortX(arr []int, left, right int) {
	//数据量级小的时候直接插入排序
	if right-left < 4 {
		SortForMerge(arr, left, right)
	} else {
		//随机找个数字
		//swap(arr, left, rand.Int()%(right-left+1)+1)
		vdata := arr[left] //坐标数,比他小 左边。大 右边
		lt := left         // arr[left+1,lt] < vdata
		gt := right + 1    // arr[gt... right] > vdata
		i := left + 1      // arr[lt+1,i] == vdata
		for i < gt {
			if arr[i] < vdata {
				swap(arr, i, lt+1) //移动到小于的地方
				lt++               //前进循环
				i++
			} else if arr[i] > vdata {
				swap(arr, i, gt-1) //移动到大于的地方
				gt--
			} else {
				i++

			}
		}
		swap(arr, left, lt)         //交换头部位置
		QuickSortX(arr, left, lt-1) //递归处理 小于那一段
		QuickSortX(arr, gt, right)  //递归处理大于那一段

	}

}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
