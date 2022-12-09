package main

import "math"

// 98 递归
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
}

// 中序遍历
func isValidBST02(root *TreeNode) bool {
	stack := []*TreeNode{}
	tmp := math.MinInt64
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp >= root.Val {
			return false
		}
		tmp = root.Val
		root = root.Right
	}
	return true
}
