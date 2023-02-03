package main

import "math"

// 515 每个树行中找最大值
// BFS
func largestValues(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for step := 0; stack != nil; step++ { // step用于记录当前层
		q := stack
		stack = nil
		if len(res) <= step {
			res = append(res, math.MinInt)
		}
		for i := 0; i < len(q); i++ {
			// 获取最大值,当前层
			res[step] = compare(q[i].Val, res[step])
			if q[i].Left != nil {
				stack = append(stack, q[i].Left)
			}
			if q[i].Right != nil {
				stack = append(stack, q[i].Right)
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
func largestValues02(root *TreeNode) (res []int) {
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
		res[level] = compare(node.Val, res[level])
		if node.Left != nil {
			// drill down
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			// drill down
			dfs(node.Right, level+1)
		}
	}
	// call dfs
	dfs(root, 0)
	return
}
