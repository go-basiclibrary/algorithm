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

// Morris
func preorderTraversal03(root *TreeNode) []int {
	res := make([]int, 0)
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left //前驱节点
		if p2 != nil {
			//找最右nil
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				//存储当前节点的数据
				res = append(res, p1.Val)
				p1 = p1.Left
				continue
			}
			p2.Right = nil //遍历完了,到了回根阶段,丢弃p1
		} else {
			// 左
			res = append(res, p1.Val)
		}
		p1 = p1.Right
	}
	return res
}
