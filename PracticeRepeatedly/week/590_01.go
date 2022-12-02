package main

type Node struct {
	Val      int
	Children []*Node
}

// 递归
func postorder(root *Node) []int {
	res := make([]int, 0)
	var porder func(node *Node)
	porder = func(node *Node) {
		if node == nil {
			return
		}
		for _, v := range node.Children {
			porder(v)
		}
		res = append(res, node.Val)
	}
	porder(root)
	return res
}

// 迭代  转换成左右根  根右左
func postorder02(root *Node) []int {
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
