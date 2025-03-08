// 20. Valid Parentheses

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:
// Input: s = "()"
// Output: true

// Example 2:
// Input: s = "()[]{}"
// Output: true

// Example 3:
// Input: s = "(]"
// Output: false

// Example 4:
// Input: s = "([])"
// Output: true

// Constraints:
// 1 <= s.length <= 10^4
// s consists of parentheses only '()[]{}'.
package main

func isStartBr(br rune) bool {
	return br == '(' || br == '[' || br == '{'
}

func getStartBr(br rune) rune {
	switch br {
	case ')':
		return '('
	case '}':
		return '{'
	case ']':
		return '['
	default:
		panic("wrong closed bracket")
	}
}

func isValid(s string) bool {
	st := make([]rune, 0, len(s)/2)
	for _, br := range s {
		if isStartBr(br) {
			st = append(st, br)
		} else {
			if len(st) == 0 {
				return false
			}
			last := st[len(st)-1]
			st = st[:len(st)-1]
			if last != getStartBr(br) {
				return false
			}
		}
	}
	return len(st) == 0
}
