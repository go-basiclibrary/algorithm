package practice

type Node struct {
	Val      int
	Children []*Node
}

// 590 N叉树的后序遍历  递归
func postorder(root *Node) []int {
	var porder func(node *Node)
	res := make([]int, 0)
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

// 迭代
func postorder02(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	res := make([]int, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		stack = append(stack, node.Children...)
	}
	reverse(res)
	return res
}

func reverse(res []int) {
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
}
