package practice

// 105 从前序与中序遍历序列构造二叉树
// 递归  想清楚如何计算清楚左右子树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil} // 前序遍历根节点为头
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] { // 找到中序遍历根,其左边为左子树,右边有右子树
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])   // 前序的左子树以及中序的左子树
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:]) //右子树
	return root
}
