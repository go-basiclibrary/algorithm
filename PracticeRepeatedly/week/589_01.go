package main

// 589 N叉树前序遍历 根左右
func preorder(root *Node) []int {
	var porder func(node *Node)
	res := make([]int, 0)
	porder = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, c := range node.Children {
			porder(c)
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
	stack := []*Node{root}
	res := make([]int, 0)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		// 需要倒叙内部才能形成右先入栈
		reverseP(root.Children)
		stack = append(stack, root.Children...)
	}
	return res
}

func reverseP(c []*Node) {
	for i := 0; i < len(c)/2; i++ {
		c[i], c[len(c)-i-1] = c[len(c)-i-1], c[i]
	}
}
