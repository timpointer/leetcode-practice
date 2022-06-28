package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// get one value from l1

	// get one value from l2

	// if v1 <= v2 or l2 is empty, then add v1 to new list, get another v1

	// if v1 > v2 or l1 is empty, then add v2 to new list, get another v2

	// util l1 and l2 is empty then stop
	list := []int{}

	v1 := list1
	v2 := list2

	for v1 != nil || v2 != nil {
		if v2 == nil {
			list = append(list, v1.Val)
			v1 = v1.Next
			continue
		}

		if v1 == nil {
			list = append(list, v2.Val)
			v2 = v2.Next
			continue
		}

		if v1.Val <= v2.Val {
			list = append(list, v1.Val)
			v1 = v1.Next
			continue
		}

		if v1.Val > v2.Val {
			list = append(list, v2.Val)
			v2 = v2.Next
			continue
		}

	}

	return NewLinkList(list)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
