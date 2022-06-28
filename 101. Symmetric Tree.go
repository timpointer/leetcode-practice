package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right != nil {
		return false
	}
	if left != nil && right == nil {
		return false
	}
	if left == nil && right == nil {
		return true
	}

	if left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left) {
		return true
	}
	return false
}
