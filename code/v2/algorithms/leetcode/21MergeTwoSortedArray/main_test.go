package main

import (
	"testing"
	"github.com/kr/pretty"
)

func TestReverse(t *testing.T) {
	test(t, mergeTwoLists)
}

func test(t *testing.T, f func(l1 *ListNode, l2 *ListNode) *ListNode) {
	tables := []struct {
		l1     *ListNode
		l2     *ListNode
		result *ListNode
	}{
		{
			&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			&ListNode{1, &ListNode{1, &ListNode{2,
				&ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
	}

	for _, table := range tables {
		result := f(table.l1, table.l2)
		for {
			if table.result == nil {
				if result != nil {
					t.Errorf("table.result == nil, resul != null. %v. Got: %v", table, result)
					break
				}
				break
			}
			if table.result.Val != result.Val {
				t.Errorf("incorrect %s. Got: %s",
					pretty.Sprint(table),
					pretty.Sprint(result))
				break
			}
			table.result = table.result.Next
			result = result.Next
		}
	}
}
