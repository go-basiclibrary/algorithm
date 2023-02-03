package main

import "math"

// 515 在每个树行中找最大值
// BFS
func largestValues(root *TreeNode) []int {
	// terminate
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	var res []int
	for floor := 0; stack != nil; floor++ {
		cur := stack
		stack = nil
		if len(res) <= floor {
			res = append(res, math.MinInt)
		}
		for i := 0; i < len(cur); i++ {
			// 比较大小
			res[floor] = compare(res[floor], cur[i].Val)
			// 收集下一层左右根
			if cur[i].Left != nil {
				stack = append(stack, cur[i].Left)
			}
			if cur[i].Right != nil {
				stack = append(stack, cur[i].Right)
			}
		}
	}
	return res
}

func compare(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS
func largestValues02(root *TreeNode) []int {
	var res []int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		// terminate
		if node == nil {
			return
		}
		// current logic
		if len(res) <= level {
			res = append(res, math.MinInt)
		}
		res[level] = compare(res[level], node.Val)
		// next floor
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	// call dfs
	dfs(root, 0)
	return res
}
