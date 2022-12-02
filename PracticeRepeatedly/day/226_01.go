package main

// 226 翻转二叉树
// 递归
func invertTree(root *TreeNode) *TreeNode {
	//终止条件
	if root == nil {
		return root
	}
	//进入到下一层
	l := invertTree(root.Left)
	r := invertTree(root.Right)
	//逻辑
	root.Left = r
	root.Right = l
	//清理当前层
	return root
}

// 迭代
func invertTree02(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := tmp.Left
		tmp.Left = tmp.Right
		tmp.Right = left
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
	}
	return root
}
