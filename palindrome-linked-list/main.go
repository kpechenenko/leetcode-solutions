// 234. Palindrome Linked List

// Given the head of a singly linked list, return true if it is a  palindrome or false otherwise.

// Example 1:
// Input: head = [1,2,2,1]
// Output: true

// Example 2:
// Input: head = [1,2]
// Output: false

// Constraints:

// The number of nodes in the list is in the range [1, 105].
// 0 <= Node.val <= 9
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head
	// find middle
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// reverse middle...end
	var reversed *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = reversed
		reversed = slow
		slow = next
	}
	// check start...middle equals to middle...end
	left, right := head, reversed
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		right = right.Next
		left = left.Next
	}
	return true
}
