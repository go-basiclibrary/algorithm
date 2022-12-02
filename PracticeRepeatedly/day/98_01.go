package main

import "math"

// 98 验证二叉搜索树
// 中序判断
func isValidBST(root *TreeNode) bool {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

// 递归
func isValidBST02(root *TreeNode) bool {
	return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return valid(root.Left, lower, root.Val) && valid(root.Right, root.Val, upper)
}

// 中序直接判断
func isValidBST03(root *TreeNode) bool {
	stack := []*TreeNode{}
	inorder := math.MinInt64
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
