package main

/*
2->4->3  245 545
5->6->4  565 565
7->0->8  807 1110

解题思路：递归
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
        // 递归终止条件
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