package main

// 前中后序遍历  迭代方式
// 前序
func preorderTraversal04(root *TreeNode) []int {
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

// 中序 左根右
func inorderTraversal(root *TreeNode) []int {
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
	return res
}

// 后序  左右根   根右左
func postorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	res := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Right
		}
		root = stack[len(stack)-1].Left
		stack = stack[:len(stack)-1]
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
