package main

// 144 前序遍历
// 递归
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		//终止条件
		if node == nil {
			return
		}
		//当前层条件(根左右)
		res = append(res, node.Val)
		//进入下一层
		preorder(node.Left)  //左
		preorder(node.Right) //右
		//清理当前层
	}
	preorder(root)

	return res
}

// 迭代
func preorderTraversal02(root *TreeNode) []int {
	//维护一个递归栈
	stack := []*TreeNode{}
	res := make([]int, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			// 根
			res = append(res, root.Val)
			stack = append(stack, root.Right)
			// 左
			root = root.Left
		}
		// 右
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return res
}
