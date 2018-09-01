// Source : https://leetcode.com/problems/add-two-numbers/description/
// Author : Lei Cao
// Date   : 2018-07-12

/**********************************************************************************
* You are given two non-empty linked lists representing two non-negative integers.
* The digits are stored in reverse order and each of their nodes contain a single digit.
* Add the two numbers and return it as a linked list.
*
* Example:
*
* Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
* Output: 7 -> 0 -> 8
* Explanation: 342 + 465 = 807.
*
**********************************************************************************/

package main

import (
	"github.com/kr/pretty"
)

func main() {
	pretty.Println(addTwoNumbers(
		&ListNode{1, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{1, &ListNode{6, &ListNode{4, nil}}},
	))

	l1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}

	pretty.Println(l1)
	pretty.Println(Reverse(l1))

	l2 := &ListNode{1, &ListNode{8, nil}}
	pretty.Println(l2)
	pretty.Println(Reverse(l2))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(l *ListNode) *ListNode {
	if l.Next == nil {
		return l
	}
	first := l
	rest := l.Next
	temp := Reverse(rest)
	first.Next.Next = first
	first.Next = nil
	return temp
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	// keep track of the last node
	rear := result
	carry := 0
	for i := 0; ; i++ {
		if l1 == nil && l2 == nil {
			if carry > 0 {
				newNode := &ListNode{Val: 1}
				rear.Next = newNode
			}
			break
		}
		if l1 != nil {
			carry += l1.Val
		}
		if l2 != nil {
			carry += l2.Val
		}
		val := carry % 10
		if carry > 9 {
			carry = 1
		} else {
			carry = 0
		}
		if i == 0 {
			result.Val = val
		} else {
			newNode := &ListNode{Val: val}
			rear.Next = newNode
			rear = newNode
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return result
}
