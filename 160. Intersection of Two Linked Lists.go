package main

// TODO have better method
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	for a != nil {
		for b != nil {
			if a == b {
				return a
			}
			b = b.Next
		}
		b = headB
		a = a.Next
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	kv := map[*ListNode]int{}
	a := headA
	b := headB
	for a != nil {
		kv[a] = 1
		a = a.Next
	}
	for b != nil {
		_, ok := kv[b]
		if ok == true {
			return b
		}
		b = b.Next
	}
	return nil
}
