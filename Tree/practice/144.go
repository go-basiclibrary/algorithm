package practice

// 144 二叉树的前序遍历
// 递归
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		//递归终结条件
		if node == nil {
			return
		}
		//当前曾逻辑
		res = append(res, node.Val)
		//下到下层(核心,根左右)
		preorder(node.Left)
		preorder(node.Right)
		//清理当前层
	}
	preorder(root)

	return res
}

// 迭代  存根求右
func preorderTraversal02(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			//根
			res = append(res, root.Val)
			stack = append(stack, root) // 存储根,后续解决右
			// 左
			root = root.Left
		}
		// 右
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return res
}
