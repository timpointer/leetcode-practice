package main

func NewLinkList(ints []int) *ListNode {
	var firstNode *ListNode
	var lastNode *ListNode
	for i, val := range ints {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		if i == 0 {
			firstNode = node
		}

		if lastNode != nil {
			lastNode.Next = node
		}
		lastNode = node
	}
	return firstNode
}

func NewIntList(firstNode *ListNode) []int {
	list := []int{}

	lastNode := firstNode
	for lastNode != nil {
		list = append(list, lastNode.Val)
		lastNode = lastNode.Next
	}

	return list
}

func ListToList(input []int) []int {
	list := NewLinkList(input)
	return NewIntList(list)
}
