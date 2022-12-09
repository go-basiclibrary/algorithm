package main

// 111 二叉树的公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//单左or右节点返回条件
	if root.Val == p.Val || root.Val == q.Val {
		return root // 返回当前节点
	}
	//找分支  如果找到左右则是当前节点为祖先
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 控制退出条件
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
