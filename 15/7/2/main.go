package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func a(root *TreeNode) (res []int) {
	var pre func(node *TreeNode)
	pre = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		pre(node.Left)
		pre(node.Right)
	}
	pre(root)
	return res
}

func b(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return res
}

func c(root *TreeNode) (res []int) {
	var pos func(node *TreeNode)
	pos = func(node *TreeNode) {
		if node == nil {
			return
		}
		pos(node.Left)
		pos(node.Right)
		res = append(res, node.Val)
	}
	pos(root)
	return res
}
