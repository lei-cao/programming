// Source : https://oj.leetcode.com/problems/merge-sorted-array/
// Author : Lei Cao
// Date   : 2018-09-01

/********************************************************************************** 
* 
* Merge two sorted linked lists and return it as a new list.
* The new list should be made by splicing together the nodes of the first two lists.
*
* Example:
*
* Input: 1->2->4, 1->3->4
* Output: 1->1->2->3->4->4
*               
**********************************************************************************/

package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = new(ListNode)
	var last = result
	for l1 != nil || l2 != nil {
		if l1 != nil {
			if l2 != nil {
				if l1.Val < l2.Val {
					last.Next = l1
					last = l1
					l1 = l1.Next
				} else if l1.Val == l2.Val {
					last.Next = l1
					last = l1
					l1 = l1.Next

					last.Next = l2
					last = l2
					l2 = l2.Next
				} else {
					last.Next = l2
					last = l2
					l2 = l2.Next
				}
			} else {
				last.Next = l1
				last = l1
				l1 = l1.Next
			}
		} else if l2 != nil {
			last.Next = l2
			last = l2
			l2 = l2.Next
		}
	}
	return  result.Next
}
