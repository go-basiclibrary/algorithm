package practice

// Codec 297 二叉树的序列化和反序列化
//type Codec struct {
//}
//
//func Constructor() (res Codec) {
//	return
//}
//
//// Serializes a tree to a single string.
//func (c *Codec) serialize(root *TreeNode) string {
//	res := ""
//	var serialized func(node *TreeNode)
//	serialized = func(node *TreeNode) {
//		if node == nil {
//			res += "null,"
//			return
//		}
//		res += strconv.Itoa(node.Val) + ","
//		serialized(node.Left)
//		serialized(node.Right)
//	}
//	serialized(root)
//	return res
//}
//
//// Deserializes your encoded data to tree.
//func (c *Codec) deserialize(data string) *TreeNode {
//	sp := strings.Split(data, ",")
//	var build func() *TreeNode
//	build = func() *TreeNode {
//		if sp[0] == "null" {
//			sp = sp[1:]
//			return nil
//		}
//		val, _ := strconv.Atoi(sp[0])
//		sp = sp[1:]
//		return &TreeNode{val, build(), build()}
//	}
//	return build()
//}

//type Codec struct {
//}
//
//func Constructor() Codec {
//	return Codec{}
//}
//
//// Serializes a tree to a single string.
//func (c *Codec) serialize(root *TreeNode) string {
//	bd := strings.Builder{}
//	var serialized func(node *TreeNode)
//	serialized = func(node *TreeNode) {
//		if node == nil {
//			bd.WriteString("null,")
//			return
//		}
//		bd.WriteString(strconv.Itoa(node.Val) + ",")
//		serialized(node.Left)
//		serialized(node.Right)
//	}
//	serialized(root)
//	return bd.String()
//}
//
//// Deserializes your encoded data to tree.
//func (c *Codec) deserialize(data string) *TreeNode {
//	sp := strings.Split(data, ",")
//	var build func() *TreeNode
//	build = func() *TreeNode {
//		if sp[0] == "null" {
//			sp = sp[1:]
//			return nil
//		}
//		s, _ := strconv.Atoi(sp[0])
//		sp = sp[1:]
//		return &TreeNode{s, build(), build()}
//	}
//	return build()
//}

// func2 层序遍历

//type Codec struct {
//}
//
//func Constructor() (_ Codec) {
//	return
//}
//
//// Serializes a tree to a single string.
//func (c *Codec) serialize(root *TreeNode) string {
//	var res []string
//	// 层序遍历
//	var stack = []*TreeNode{root}
//	for len(stack) > 0 {
//		root = stack[0]
//		stack = stack[1:]
//		if root != nil {
//			res = append(res, strconv.Itoa(root.Val))
//			stack = append(stack, root.Left)
//			stack = append(stack, root.Right)
//		} else {
//			res = append(res, "null")
//		}
//	}
//	return strings.Join(res, ",")
//}
//
//// Deserializes your encoded data to tree.
//func (c *Codec) deserialize(data string) *TreeNode {
//	if data == "null" {
//		return nil
//	}
//
//	// exchange list
//	list := strings.Split(data, ",")
//	val, _ := strconv.Atoi(list[0])
//	// create root node
//	var root = &TreeNode{Val: val}
//	q := []*TreeNode{root}
//	cursor := 1
//	for cursor < len(list) {
//		node := q[0]
//		q = q[1:]
//		leftVal := list[cursor]
//		rightVal := list[cursor+1]
//		if leftVal != "null" {
//			val, _ = strconv.Atoi(leftVal)
//			node.Left = &TreeNode{Val: val}
//			q = append(q, node.Left)
//		}
//		if rightVal != "null" {
//			val, _ = strconv.Atoi(rightVal)
//			node.Right = &TreeNode{Val: val}
//			q = append(q, node.Right)
//		}
//		cursor += 2
//	}
//	return root
//}
