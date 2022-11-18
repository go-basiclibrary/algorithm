package practice

// TreeNode 94 二叉树的中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		//终结条件
		if node == nil {
			return
		}
		//处理逻辑
		inorder(node.Left)
		res = append(res, node.Val)
		//下到树层
		inorder(node.Right)
		//清理当前层
	}
	inorder(root)
	return res
}

// 迭代
func inorderTraversal02(root *TreeNode) []int {
	// 思想:先拿出所有的左节点,依次用一个栈维护,然后再基于此遍历左栈
	res := make([]int, 0)
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		// 维护所有左树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 再次回到循环开始执行stack
		// 根据stack的左子树,求完求一下右子树
		root = stack[len(stack)-1]
		// 出栈
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
