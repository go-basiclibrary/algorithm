package practice

// 429 N叉树的层序遍历
// 递归
func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	res := make([][]int, 0)
	for stack != nil {
		level := []int{}
		temp := stack
		stack = nil
		for _, node := range temp {
			level = append(level, node.Val)
			stack = append(stack, node.Children...)
		}
		res = append(res, level)
	}
	return res
}
