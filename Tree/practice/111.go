package practice

// 111 二叉树最小深度
//深度搜索
func minDepth(root *TreeNode) int {
	var mDepth func(node *TreeNode) int
	mDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Right == nil && node.Left != nil {
			return minDepth(root.Right) + 1
		}
		if node.Right != nil && node.Left == nil {
			return minDepth(root.Left) + 1
		}
		return min(mDepth(node.Left), mDepth(node.Right)) + 1
	}
	return mDepth(root)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
