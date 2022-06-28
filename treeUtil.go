package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(list []int) *TreeNode {
	return NewTreeHelper(list, 0, 0)
}

func NewTreeHelper(list []int, level int, idx int) *TreeNode {
	maxNum := len(list) - 1
	leftIdx := idx*level + 1
	rightIdx := idx*level + 2

	if leftIdx > maxNum {
		return nil
	}

	if rightIdx > maxNum {
		return nil
	}
	left := NewTreeHelper(list, 0, leftIdx)
	right := NewTreeHelper(list, 0, rightIdx)

	val := list[idx]
	root := &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}

	return root
}

func NewTreeList(root *TreeNode) []int {
	return inorderTraversal(root)
}
