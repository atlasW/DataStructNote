package main

import "fmt"

// 凑零钱问题
// 给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1 。

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
	fmt.Println(coinChange4([]int{1, 2, 5}, 16))
}

//递归 暴力破解
func coinChange2(coins []int, amount int) (n int) {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res1 := func(n int) int {
		res := n + 1
		for _, coin := range coins {
			subproblem := coinChange(coins, n-coin)
			if subproblem >= 0 {
				res = min(res, subproblem) + 1
			}
		}
		return res
	}(amount)
	//求最小值  设置一个极大值
	return res1

}

//带备忘录的递归
func coinChange3(coins []int, amount int) (n int) {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	db := make([]int, amount+1, amount+1)
	var subproblem int
	res1 := func(n int) int {
		res := n + 1
		for _, coin := range coins {
			//备忘录处理
			if db[n-coin] == 0 {
				subproblem = coinChange(coins, n-coin)
				db[n-coin] = subproblem
			} else {
				subproblem = db[n-coin]
			}
			if subproblem >= 0 {
				res = min(res, subproblem) + 1
			}
		}
		return res
	}(amount)
	//求最小值  设置一个极大值
	return res1

}

//dp数组 迭代
func coinChange4(coins []int, amount int) (n int) {
	dp := make([]int, amount+1, amount+1)
	dp[0] = 0
	//构建dp表    迭代
	/////this
	for k := 1; k <= amount; k++ {
		for _, coin := range coins {
			if k-coin < 0 {
				continue
			} else {
				if dp[k] != 0 {
					dp[k] = min(dp[k], dp[k-coin]+1)
				} else {
					dp[k] = dp[k-coin] + 1
				}

			}
		}
	}
	fmt.Println(dp)
	return dp[amount]
}

/*
 // 外层 for 循环在遍历所有状态的所有取值
    for (int i = 0; i < dp.size(); i++) {
        // 内层 for 循环在求所有选择的最小值
        for (int coin : coins) {
            // 子问题无解，跳过
            if (i - coin < 0) continue;
            dp[i] = min(dp[i], 1 + dp[i - coin]);
        }
    }
    return (dp[amount] == amount + 1) ? -1 : dp[amount];
*/
