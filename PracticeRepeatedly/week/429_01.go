package main

// 层序遍历
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	stack := []*Node{root}
	for stack != nil {
		level := []int{}
		tmp := stack
		stack = nil
		for _, v := range tmp {
			level = append(level, v.Val)
			stack = append(stack, v.Children...)
		}
		res = append(res, level)
	}
	return res
}
