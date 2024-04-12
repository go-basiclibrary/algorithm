package main

// 98 验证二叉搜索树
func main() {
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// func1 递归
//func isValidBST(root *TreeNode) bool {
//	return helper(root, math.MinInt64, math.MaxInt64)
//}
//
//func helper(root *TreeNode, lower, upper int) bool {
//	if root == nil {
//		return true
//	}
//	if root.Val <= lower || root.Val >= upper {
//		return false
//	}
//	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
//}

// func2 中序遍历 stack
//func isValidBST(root *TreeNode) bool {
//	var stack []*TreeNode
//	var inorder = math.MinInt64
//	for root != nil || len(stack) > 0 {
//		for root != nil {
//			stack = append(stack, root)
//			root = root.Left
//		}
//		root = stack[len(stack)-1]
//		stack = stack[:len(stack)-1]
//		if root.Val <= inorder {
//			return false
//		}
//		inorder = root.Val
//		root = root.Right
//	}
//	return true
//}
