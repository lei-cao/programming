package main

import (
	"testing"
	"github.com/kr/pretty"
)

func TestAddTwoNumbers(t *testing.T) {
	test(t, addTwoNumbers)
}

func test(t *testing.T, f func(l1 *ListNode, l2 *ListNode) *ListNode) {
	tables := []struct {
		l1 *ListNode
		l2 *ListNode
		result *ListNode
	}{
		{&ListNode{1, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{1, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{2, &ListNode{0, &ListNode{8, nil}}},
		},
		{&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{&ListNode{1, &ListNode{8, nil}},
			&ListNode{0, nil},
			&ListNode{1, &ListNode{8, nil}},
		},
		{&ListNode{5, nil},
			&ListNode{5, nil},
			&ListNode{0, &ListNode{1, nil}},
		},
	}

	for _, table := range tables {
		result := f(table.l1, table.l2)
		for {
			if table.result == nil {
				if result != nil {
					t.Errorf("incorrect %v. Got: %v", table, result)
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
