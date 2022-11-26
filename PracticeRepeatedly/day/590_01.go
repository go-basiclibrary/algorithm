package main

type Node struct {
	Val      int
	Children []*Node
}

// 590 N叉树 后序遍历
// 递归
func postorder(root *Node) []int {
	res := make([]int, 0)
	var porder func(node *Node)
	porder = func(node *Node) {
		if node == nil {
			return
		}
		for _, c := range node.Children {
			porder(c)
		}
		res = append(res, node.Val)
	}
	porder(root)
	return res
}

// 迭代 左右根  根右左
func postorder02(root *Node) []int {
	res := make([]int, 0)
	stack := []*Node{}
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		stack = append(stack, root.Children...)
	}
	reverseForOrder(res)
	return res
}

func reverseForOrder(res []int) {
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
}
