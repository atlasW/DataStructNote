package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

/*
顺序搜索   遍历
*/
const N = 17186792

type Location struct {
	ID string    `json:"_id"`
	V  int       `json:"v"`
	P  int       `json:"p"`
	I  string    `json:"i"`
	O  float64   `json:"o"`
	S  int       `json:"s"`
	L  []float64 `json:"l"`
	T  struct {
		Date time.Time `json:"$date"`
	} `json:"t"`
	D int `json:"d"`
	B int `json:"b"`
	C struct {
		Date time.Time `json:"$date"`
	} `json:"c"`
}

func main() {
	starttime := time.Now()
	path := "../example/2.json"
	jsonfile, _ := os.Open(path)
	i := 0 //行数
	br := bufio.NewReader(jsonfile)
	StructArry := make([]Location, N, N)
	var line []byte
	var err error
	var value Location
	for {
		line, _, err = br.ReadLine()
		if err == io.EOF {
			break
		}
		json.Unmarshal(line, &value)
		StructArry[i] = value
		i++
	}
	fmt.Println("载入数据用时:", time.Since(starttime))
	QuickSort(StructArry, 0, len(StructArry)-1)
	fmt.Println("排序成功。")
	//
	for {
		var a int
		var b []Location
		fmt.Printf("请输入需要查询的s:")
		fmt.Scanf("%d", &a)
		starttime := time.Now()
		index := BinSearch(StructArry, a)
		fmt.Println(index)
		for i := index; ; i++ {
			b = append(b, StructArry[i])
			if StructArry[i].S != a {
				break
			}

		}
		fmt.Println("本次查询用了：", time.Since(starttime))
		fmt.Println(len(b))
		b = []Location{}
	}
}

//快速排序

func partition(a []Location, left, right, pivotIndex int) int {
	//基准值
	pivotValue := a[pivotIndex].S
	// 基准和最后一位 互换
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	storeIndex := left
	//找到基准所在的位置   //很巧妙的方法  找到基准的位置  并且将 小于基准的移动到基准的左边
	for i := left; i < right; i++ {
		if a[i].S <= pivotValue {
			a[storeIndex], a[i] = a[i], a[storeIndex]
			storeIndex++
		}

	}
	a[right], a[storeIndex] = a[storeIndex], a[right]
	return storeIndex

}
func QuickSort(a []Location, left, right int) {
	if right > left {
		pivotIndex := left
		pivotNewIndex := partition(a, left, right, pivotIndex)
		QuickSort(a, left, pivotNewIndex-1)
		QuickSort(a, pivotNewIndex+1, right)
	}
}

//二分查找
func BinSearch(arr []Location, data int) int {
	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2
		if arr[mid].S > data {
			right = mid - 1
		} else if arr[mid].S < data {
			left = mid + 1
		} else {
			return mid
		}

	}
	return -1
}
