package main

// 111 二叉树的最小深度
// 深度优先
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil && root.Left != nil {
		return minDepth(root.Left) + 1
	}
	return minD(minDepth(root.Left), minDepth(root.Right)) + 1
}

func minD(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// 广度优先
func minDepth02(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{root}
	count := []int{}
	count = append(count, 1)
	for i := 0; i < len(stack); i++ {
		node := stack[i]
		depth := count[i]
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			count = append(count, depth+1) // 记录每一层的高度
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			count = append(count, depth+1)
		}
	}
	return 0
}
