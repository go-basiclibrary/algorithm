package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144 二叉树前序遍历 根左右
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	//递归
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		//终止条件
		if node == nil {
			return
		}
		//逻辑
		res = append(res, node.Val)
		//下层
		preorder(node.Left)
		preorder(node.Right)
		//清理当前层
	}
	preorder(root)

	return res
}

// 迭代
func preorderTraversal02(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
