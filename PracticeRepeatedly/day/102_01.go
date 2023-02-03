package main

// 102 二叉树的层序遍历
// BFS
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for i := 0; len(stack) > 0; i++ {
		cur := []*TreeNode{}
		var now []int
		for j := 0; j < len(stack); j++ {
			// 获取每一层的数据
			now = append(now, stack[j].Val)
			if stack[j].Left != nil {
				cur = append(cur, stack[j].Left)
			}
			if stack[j].Right != nil {
				cur = append(cur, stack[j].Right)
			}
		}
		res = append(res, now)
		stack = cur
	}
	return res
}

// DFS
func levelOrder02(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		// terminate
		if node == nil {
			return
		}
		// current logic
		if len(res) <= level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		if node.Left != nil {
			// Next floor
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			// Next floor
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	return res
}
