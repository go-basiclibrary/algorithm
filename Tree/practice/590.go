package practice

// 二叉树的 后续遍历
// 递归
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		//退出条件
		if node == nil {
			return
		}
		//进入下层  左右根
		postorder(node.Left)
		postorder(node.Right)
		//当前层逻辑
		res = append(res, node.Val)
		//清理当前层
	}
	postorder(root)
	return res
}

// 迭代
func postorderTraversal02(root *TreeNode) (res []int) {
	stack := []*TreeNode{} //栈
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return res
}
