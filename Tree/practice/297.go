package practice

import (
	"strconv"
	"strings"
)

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

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	bd := strings.Builder{}
	var serialized func(node *TreeNode)
	serialized = func(node *TreeNode) {
		if node == nil {
			bd.WriteString("null,")
			return
		}
		bd.WriteString(strconv.Itoa(node.Val) + ",")
		serialized(node.Left)
		serialized(node.Right)
	}
	serialized(root)
	return bd.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		s, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{s, build(), build()}
	}
	return build()
}
