package main

// 226 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	var ivert func(node *TreeNode)
	ivert = func(node *TreeNode) {
		if node == nil {
			return
		}
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
		ivert(node.Left)
		ivert(node.Right)
	}
	ivert(root)
	return root
}

// 迭代
func invertTree02(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := tmp.Left
		tmp.Left = tmp.Right
		tmp.Right = left
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}
	}
	return root
}
