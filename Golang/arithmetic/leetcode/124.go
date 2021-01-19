package main

import (
	"fmt"
	"math"
)

// 二叉树忠的最大路径和
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//  -10   9  20  _  _   15  7
func CreatTree() (tree *TreeNode) {
	a := &TreeNode{9, nil, nil}
	c := &TreeNode{15, nil, nil}
	d := &TreeNode{7, nil, nil}
	b := &TreeNode{20, c, d}
	tree = &TreeNode{-10, a, b}
	return
}

//
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := max(0, dfs(node.Left)), max(0, dfs(node.Right)) // 0参与取大消除负数带来的副作用
		ans = max(ans, node.Val+l+r)                            // 左+右+根最大, 获取当前节点所有最大和路径
		return max(l, r) + node.Val                             // 左+根，右+根 取最大
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := CreatTree()
	fmt.Println(maxPathSum(a))
}
