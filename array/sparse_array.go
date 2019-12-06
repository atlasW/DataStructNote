package main

/*
稀疏数组，压缩数组节省数组存储占用空间 //后改用slice适配所有规格slice
本例为二维数组
*/
import (
	"errors"
	"fmt"
)

//参数 二维数组的行数和列数，还有最常出现的值
func Sparse(orginal [][]int, often int) (sparsed [][]int) {
	high := len(orginal)
	wide := len(orginal[0])
	sparsed = append(sparsed, []int{wide, high, often})
	for row, v := range orginal {
		for column, v1 := range v {
			if v1 != often {
				sparsed = append(sparsed, []int{row, column, v1})
			}

		}
	}
	return
}

func SparseRestore(sparsed [][]int) (result [][]int, err error) {
	if len(sparsed[0]) != 3 {
		err = errors.New("不符合二维数组稀疏数组。")
		return
	}
	// restore
	//first restore with all often
	totalRow := sparsed[0][0]
	totalCol := sparsed[0][1]
	often := sparsed[0][2]
	var temp []int
	for i := 0; i < totalCol; i++ {
		for j := 0; j < totalRow; j++ {
			temp = append(temp, often)
		}
		result = append(result, temp)
		temp = []int{}
	}
	//填充数据
	for k, v := range sparsed {
		if k != 0 {
			result[v[0]][v[1]] = v[2]
		}
	}
	return
}

func main() {
	//初始化原始数据  high 20 ,wide 10
	bb := [][]int{}
	for i := 0; i < 15; i++ { // high
		aa := []int{}
		for j := 0; j < 10; j++ { //wide
			aa = append(aa, 0)
		}
		bb = append(bb, aa)
	}

	OriArray := bb
	OriArray[3][4] = 1
	OriArray[5][5] = 2
	OriArray[3][3] = 1
	fmt.Println("原始数组：")
	for _, i := range OriArray {
		fmt.Println(i)
	}
	fmt.Println("压缩后数据:")
	fmt.Println(Sparse(OriArray, 0))
	fmt.Println("恢复后的数据")
	a, _ := SparseRestore(Sparse(OriArray, 0))
	for _, i := range a {
		fmt.Println(i)
	}
}
