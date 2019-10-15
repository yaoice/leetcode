package main

import "fmt"

/*
2->4->3  243
5->6->4  564
7->0->8  807

递归
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    temp := result
    v, n := 0, 0

    for {
        v, n = add(l1, l2, n)
        temp.Val = v
        l1, l2 = next(l1), next(l2)
        if l1 == nil && l2 == nil {
            break
        }
        temp.Next = &ListNode{}
        temp = temp.Next
    }
    if n == 1 {
        temp.Next = &ListNode{Val: 1}
    }
    return result
}

// next 进入l的下一位。
func next(l *ListNode) *ListNode {
    if l != nil {
        return l.Next
    }
    return nil
}

func add(n1, n2 *ListNode, i int) (v, n int) {
    if n1 != nil {
        v = v + n1.Val
    }
    if n2 != nil {
        v = v + n2.Val
    }

    v = v + i

    if v >= 10 {
        v = v - 10
        n = 1
    }
    return v, n
}

func main() {
    l1 := &ListNode{
        Val:  3,
        Next: &ListNode{
            Val:  4,
            Next: &ListNode{
                Val:  2,
                Next: nil,
            },
        },
    }

    l2 := &ListNode{
        Val:  4,
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
