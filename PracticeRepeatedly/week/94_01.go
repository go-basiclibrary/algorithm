package main

// 中序遍历 morris算法
func inorderTraversal02(root *TreeNode) []int {
	res := make([]int, 0)
	for root != nil {
		if root.Left != nil {
			process := root.Left
			for process.Right != nil && process.Right != root {
				process = process.Right
			}
			if process.Right == nil {
				process.Right = root
				root = root.Left
			} else {
				res = append(res, root.Val)
				process.Right = nil
				root = root.Right
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}
