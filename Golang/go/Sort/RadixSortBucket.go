package main

import "fmt"

//分布式排序   计数排序 --> 桶排序 -->基数排序
//               数的映射    函数的隐射
//计数排序
func countingSort(arr []int, minvalue, maxValue int) []int {
	bucketLen := maxValue - minvalue + 1
	bucket := make([]int, bucketLen)
	//放到桶里面
	for _, v := range arr {
		bucket[v-minvalue]++
	}
	//桶里面拿出来
	result := make([]int, len(arr))
	index := 0
	for k, v := range bucket {
		kk := k + minvalue
		for j := v; j > 0; j-- {
			result[index] = kk
			index++
		}
	}
	return result

}

//十分位桶排序  map
func BucketSort(arr []int, maxValue int) []int {
	bucketLen := (maxValue / 100) + 1
	bucket := make([][]int, bucketLen)
	// 放到桶里
	for _, v := range arr {
		//函数
		vv := v / 100
		mod := v % 100
		bucket[vv] = append(bucket[vv], mod)
	}
	fmt.Println(bucket)
	//拿出来
	//	[[1,3,5],[5,6,7]]
	result := make([]int, 0)
	for k, v := range bucket {
		kk := k * 100
		SelectSort(v)
		for j := 0; j < len(v); j++ {
			result = append(result, kk+v[j])
		}
	}
	return result

}

func RadixSortBucket(arr []int) {

}
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

func main() {
	a := []int{11, 91, 222, 878, 348, 7123, 4123, 6232, 5123, 1011}
	//	a := []int{3, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 3, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9}
	fmt.Println(len(a))
	// 	b := countingSort(a, 2, 9)
	b := BucketSort(a, 7123)
	fmt.Printf("%+v\n", b)
}
