package main

import "fmt"

// 111 二叉树最小深度
func main() {
	var stack []int
	stack = append(stack, 1)
	for i := 0; i < len(stack); i++ {
		fmt.Println(len(stack))
		val := stack[i]
		fmt.Println(val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func1 递归
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	if root.Left == nil && root.Right == nil {
//		return 1
//	}
//	var minD = math.MaxInt32
//	if root.Left != nil {
//		minD = min(minDepth(root.Left), minD)
//	}
//	if root.Right != nil {
//		minD = min(minDepth(root.Right), minD)
//	}
//	return minD + 1
//}
//
//func min(depth int, depth2 int) int {
//	if depth > depth2 {
//		return depth2
//	}
//	return depth
//}

// func2 BFS
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var stack = []*TreeNode{root}
//	var count = []int{1}
//	// 遍历stack
//	for i := 0; i < len(stack); i++ {
//		node := stack[i]
//		depth := count[i]
//		if node.Left == nil && node.Right == nil {
//			return depth
//		}
//		if node.Left != nil {
//			stack = append(stack, node.Left)
//			count = append(count, depth+1)
//		}
//		if node.Right != nil {
//			stack = append(stack, node.Right)
//			count = append(count, depth+1)
//		}
//	}
//	return 0
//}

// func3 DFS
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var left = minDepth(root.Left)
//	var right = minDepth(root.Right)
//	if left == 0 || right == 0 {
//		return max(left, right) + 1
//	}
//	return min(left, right) + 1
//}
//
//func min(left int, right int) int {
//	if left > right {
//		return right
//	}
//	return left
//}
//
//func max(left int, right int) int {
//	if left > right {
//		return left
//	}
//	return right
//}
