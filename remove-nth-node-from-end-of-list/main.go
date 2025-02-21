// 19. Remove Nth Node From End of List

// Given the head of a linked list, remove the nth node from the end of the list and return its head.

// Example 1:
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]

// Example 2:
// Input: head = [1], n = 1
// Output: []

// Example 3:
// Input: head = [1,2], n = 1
// Output: [1]

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	f := head
	// make window of size n between s and f pointers
	for ; n > 0; n-- {
		f = f.Next
	}
	s := &ListNode{Val: -1, Next: head}
	dummy := s
	// move f to end of list
	for f != nil {
		f = f.Next
		s = s.Next
	}
	// delete node
	s.Next = s.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEndIterative(head *ListNode, n int) *ListNode {
	size := 0
	f := head
	for f != nil {
		size++
		f = f.Next
	}
	dummy := &ListNode{Val: -1, Next: head}
	p := dummy
	for i := 0; i < size-n; i++ {
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
