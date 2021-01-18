package main

import "fmt"

//给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。

//最少多少颗
//暴力递归
/*
 f(coins,n)     if n = 0       0
                if n <0        -1
				if n>0        min(f(coins,n-coin)) + 1
*/
func coinChange(coins []int, amount int) (n int) {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	//求最小值  设置一个极大值
	res := amount + 1
	for _, coin := range coins {
		subproblem := coinChange(coins, amount-coin)
		//递归到0  amount-coint =0   也得加上一次
		if subproblem >= 0 {
			res = min(res, subproblem) + 1
		}
	}
	return res

}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 9))
}
