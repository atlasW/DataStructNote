package main

/*
稀疏数组，压缩数组节省数组存储占用空间
本例为二维数组
*/
import (
	"errors"
	"fmt"
)

//参数 二维数组的行数和列数，还有最常出现的值
func Sparse(orginal [10][10]int, often int) (sparsed [][]int) {
	sparsed = append(sparsed, []int{cap(orginal), cap(orginal[0]), often})
	for row, v := range orginal {
		for column, v1 := range v {
			if v1 != often {
				sparsed = append(sparsed, []int{row, column, v1})
			}

		}
	}
	return
}

func SparseRestore(sparsed [][]int) (err error) {
	if len(sparsed[0]) != 3 {
		err = errors.New("不符合二维数组稀疏数组。")
		return
	}
	// restore
	totalRow = sparsed[0][0]
	totalCol := sparsed[0][1]
	often := sparsed[0][2]
	var result [totalCol][totalRow]int

	for kk, vv := range result {
		for kkk, vvv := range vv {
			result[kk][kkk] = often
		}
	}
	return
}

func main() {
	//原始数组
	var OriArray [10][10]int
	OriArray[3][4] = 1
	OriArray[5][5] = 2
	OriArray[3][3] = 1
	fmt.Println("原始数组：")
	for _, i := range OriArray {
		fmt.Println(i)
	}
	fmt.Println(Sparse(OriArray, 0))
}
