package main

// 590 N叉树后序遍历
// 递归
func postorder03(root *Node) []int {
	res := make([]int, 0)
	var porder func(node *Node)
	porder = func(node *Node) {
		//结束条件
		if node == nil {
			return
		}
		//进入下一层
		for _, c := range node.Children {
			porder(c)
		}
		//逻辑
		res = append(res, node.Val)
		//清理当前层
	}
	porder(root)
	return res
}

// 迭代  根右左
func postorder04(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	res := make([]int, 0)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		stack = append(stack, root.Children...)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
