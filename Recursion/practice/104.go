package main

// 104 二叉树的最大深度

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// func1 递归  想最小子问题,最大高度,就是左比较右高度再每次+1
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
//}
//
//func max(depth int, depth2 int) int {
//	if depth2 > depth {
//		return depth2
//	}
//	return depth
//}

// func2 广度优先 BFS
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var stack = []*TreeNode{root}
//	var depth int
//	var node *TreeNode
//	for len(stack) > 0 {
//		curStackLen := len(stack)
//		for curStackLen > 0 {
//			node = stack[0]
//			stack = stack[1:]
//			if node.Left != nil {
//				stack = append(stack, node.Left)
//			}
//			if node.Right != nil {
//				stack = append(stack, node.Right)
//			}
//			curStackLen--
//		}
//		depth++
//	}
//	return depth
//}
