package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  6,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	c := addTwoNumbers(l1, l2)
	for {
		fmt.Println(c.Val)
		if c.Next == nil {
			break
		}
		c = c.Next
	}
}
