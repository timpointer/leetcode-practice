package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TODO Follow up: Recursive solution is trivial, could you do it iteratively?
func inorderTraversal(root *TreeNode) []int {
	list := []int{}
	list = inorderTraversalHelper(root, list)
	return list
}

func inorderTraversalHelper(root *TreeNode, list []int) []int {
	if root == nil {
		return []int{}
	}
	if root.Left != nil {
		list = append(list, inorderTraversal(root.Left)...)
	}
	list = append(list, root.Val)
	if root.Right != nil {
		list = append(list, inorderTraversal(root.Right)...)
	}

	return list
}
