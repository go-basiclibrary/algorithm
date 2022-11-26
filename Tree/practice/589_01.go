package practice

// 589 N叉树的前序遍历
func preorder(root *Node) []int {
	res := make([]int, 0)
	var porder func(node *Node)
	porder = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, v := range node.Children {
			porder(v)
		}
	}
	porder(root)
	return res
}

// 迭代
func preorder02(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := []*Node{}
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// 这里要reverse一下Children
		reverseC(root.Children)
		stack = append(stack, root.Children...)
	}
	return res
}

func reverseC(c []*Node) {
	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-i-1] = c[len(c)-i-1], c[i]
	}
}
