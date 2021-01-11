package sort

import (
	"fmt"
	"math"
)

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

//分位桶排序  map
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

func RadixSort(theArray []int) []int {
	//获取最大值vl
	vl := 0
	for _, v := range theArray {
		if v > vl {
			vl = v
		}
	}
	//获取最大值的位数
	var count int = 0
	for vl%10 > 0 {
		vl = vl / 10
		count++
	}

	//给桶中对应的位置放数据
	for i := 0; i < count; i++ {
		fmt.Println(theArray)
		theData := int(math.Pow10(i)) //10的i次方
		//建立空桶
		var bucket [10][10]int
		for k := 0; k < len(theArray); k++ {
			theResidue := (theArray[k] / theData) % 10 //取余
			var childArray [10]int                     //= bucket[theResidue];//获取子数组
			for m := 0; m < 10; m++ {
				if bucket[theResidue][m] == 0 {
					childArray[m] = theArray[k]
					bucket[theResidue][m] = childArray[m]
					break
				} else {
					continue
				}
			}
		}
		//一遍循环完之后需要把数组二维数据进行重新排序，比如数组开始是10 1 18 30 23 12 7 5 18 233 144 ，循环个位数
		//循环之后的结果为10 30 1 12 23 233 144 5 7 18 18 ，然后循环十位数，结果为1 5 7 10 12 18 18 23 30 233 144
		//最后循环百位数，结果为1 5 7 10 12 18 18 23 30 144 233
		var x = 0
		slice := make([]int, len(theArray))
		for p := 0; p < len(bucket); p++ {
			for q := 0; q < len(bucket[p]); q++ {
				if bucket[p][q] != 0 {
					slice[x] = bucket[p][q]
					x++
				} else {
					break
				}
			}
		}

		for key, value := range slice {
			theArray[key] = value
		}
	}
	return theArray

}

//func SelectSort(this []int) {
//	lenth := len(this)
//	if lenth <= 1 {
//	} else {
//		//只剩一个元素 不需要挑选
//		for i := 0; i < lenth-1; i++ {
//			min := i
//			for j := i + 1; j < lenth; j++ {
//				if this[min] > this[j] {
//					min = j
//				}
//			}
//			if i != min {
//				this[i], this[min] = this[min], this[i]
//			}
//		}
//	}
//}
//
//func main() {
//	a := []int{11, 91, 222, 878, 348, 7123, 4123, 6232, 5123, 1011}
//	//	a := []int{3, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 3, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9, 9, 2, 2, 5, 4, 3, 6, 2, 3, 4, 4, 9}
//	fmt.Println(len(a))
//	// 	b := countingSort(a, 2, 9)
//	//b := BucketSort(a, 7123)
//	b := RadixSort(a)
//	fmt.Printf("%+v\n", b)
//}
