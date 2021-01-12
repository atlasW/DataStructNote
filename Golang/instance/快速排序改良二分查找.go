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
	//a := []int{3, 9, 2, 8, 1, 7, 4, 6, 5, 10}
	//QuickSortPlus(a)
	//fmt.Println(a)
	starttime := time.Now()
	path := "./example/1.json"
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
	starttime = time.Now()
	QuickSortPlus(StructArry)
	fmt.Println("排序数据用时:", time.Since(starttime))
	//
	for {
		var a int
		var b []Location
		fmt.Printf("请输入需要查询的s:")
		fmt.Scanf("%d", &a)
		starttime := time.Now()
		index := BinSearch(StructArry, a)
		if index != -1 {
			b = append(b, StructArry[index])
		}
		for left := index - 1; left > 0; left-- {
			if StructArry[left].S == a {
				b = append(b, StructArry[left])
			} else {
				break
			}
		}
		for right := index + 1; right < len(StructArry)-1; right++ {
			if StructArry[right].S == a {
				b = append(b, StructArry[right])
			} else {
				break
			}
		}

		fmt.Println("本次查询用了：", time.Since(starttime))
		fmt.Println(len(b))
		b = []Location{}
	}
}

//二分查找
func BinSearch(arr []Location, data int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		fmt.Println(left)
		fmt.Println(right)
		fmt.Println("-0-----------")
		mid := (left + right) / 2
		if arr[mid].S > data {
			right = mid - 1
		} else if arr[mid].S < data {
			left = mid + 1
		} else {
			return mid
		}
	}
	fmt.Println("this")
	return -1
}

//插入排序
func SortForMerge(arr []Location, left, right int) {
	for i := left; i <= right; i++ {
		temp := arr[i] //备份数据
		var j int
		for j = i; j > left && arr[j-1].S > temp.S; j-- { //定位
			arr[j] = arr[j-1] //数据后移
		}
		arr[j] = temp
	}

}

//改良版本的快速排序  解约内存
func QuickSortPlus(arr []Location) {
	QuickSortX(arr, 0, len(arr)-1)
}

//递归快速排序
func QuickSortX(arr []Location, left, right int) {
	//数据量级小的时候直接插入排序
	if right-left < 3 {
		SortForMerge(arr, left, right)
	} else {
		//随机找个数字
		vdata := arr[left].S //坐标数,比他小 左边。大 右边
		lt := left           // arr[left+1,lt] < vdata
		gt := right + 1      // arr[gt... right] > vdata
		i := left + 1        // arr[lt+1,i] == vdata
		for i < gt {
			if arr[i].S < vdata {
				swap(arr, i, lt+1) //移动到小于的地方
				lt++               //前进循环
				i++
			} else if arr[i].S > vdata {
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

func swap(arr []Location, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
