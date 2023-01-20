package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 144 前序遍历morris算法
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left // 先找左子的右根,剪枝接入根
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				//剪枝接入
				p2.Right = p1
				res = append(res, p1.Val) // 这里收集根数据
				p1 = p1.Left              // 进一步找下一个左子的右根
				continue
			}
			p2.Right = nil
		} else {
			res = append(res, p1.Val) // 左无,要接入根
		}
		p1 = p1.Right
	}
	return res
}
