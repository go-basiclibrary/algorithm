package main

// 589 N叉树前序遍历  根左右
// 迭代
func preorder(root *Node) []int {
	res := make([]int, 0)
	var porder func(node *Node)
	porder = func(node *Node) {
		//结束条件
		if node == nil {
			return
		}
		//当前层逻辑
		res = append(res, node.Val)
		//进入下层
		for _, c := range node.Children {
			porder(c)
		}
		//清理当前层
	}
	porder(root)
	return res
}

// 迭代
func preorder02(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	res := make([]int, 0)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		//排序子node,因为根左右,入栈要遵循,右左出栈才是左右
		reverse589(root.Children)
		stack = append(stack, root.Children...)
	}

	return res
}

func reverse589(c []*Node) {
	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-i-1] = c[len(c)-i-1], c[i]
	}
}
