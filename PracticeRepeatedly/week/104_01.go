package main

// 104 二叉树深度
// 深度优先算法
func maxDepth(root *TreeNode) int {
	var mDepth func(node *TreeNode) int
	mDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return max1(mDepth(node.Left), mDepth(node.Right)) + 1
	}
	return mDepth(root)
}

func max1(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// 广度优先算法
func maxDepth02(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		sz := len(stack)
		for sz > 0 {
			node := stack[0]
			stack = stack[1:]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
