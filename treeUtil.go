package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTree(list []int) *TreeNode {
	return NewTreeHelper(list, 1, 0)
}

func NewTreeHelper(list []int, level int, idx int) *TreeNode {
	maxNum := len(list) - 1
	leftIdx := idx*level + 1
	rightIdx := idx*level + 2

	var left, right *TreeNode
	if leftIdx > maxNum {
		left = nil
	} else {
		left = NewTreeHelper(list, level+1, leftIdx)
	}

	if rightIdx > maxNum {
		right = nil
	} else {
		right = NewTreeHelper(list, level+1, rightIdx)
	}

	val := list[idx]
	if val == -1 {
		return nil
	}
	root := &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}

	return root
}

func TreeToTree(list []int) []int {
	return inorderTraversal(NewTree(list))
}
