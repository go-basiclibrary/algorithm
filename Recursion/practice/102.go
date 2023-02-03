package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 102 二叉树层序遍历  BFS实现
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	stack := []*TreeNode{root} // 模拟优先队列,先进先出
	for i := 0; len(stack) > 0; i++ {
		stackNext := []*TreeNode{}
		res = append(res, []int{})
		for j := 0; j < len(stack); j++ {
			// 加入当前层
			node := stack[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil {
				stackNext = append(stackNext, node.Left)
			}
			if node.Right != nil {
				stackNext = append(stackNext, node.Right)
			}
		}
		stack = stackNext
	}
	return res
}

// DFS实现
func levelOrder02(root *TreeNode) [][]int {
	var res [][]int
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		// terminate
		if node == nil {
			return
		}
		// cur
		if level >= len(res) {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	return res
}
