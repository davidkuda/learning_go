package main

import "fmt"


type ListNode struct {
    Val int
    Next *ListNode
}


func main() {
    l1 := ListNode{
        Val: 3,
        Next: nil,
    }

    l2 := ListNode{
        Val: 42,
        Next: &l1,
    }

    l3 := ListNode{
        Val: 108,
        Next: &l2,
    }

    cur := &l3

    for cur != nil {
        fmt.Println(cur.Val)
        cur = cur.Next
    }
}

func reverseList(head *ListNode) *ListNode {

    var tail *ListNode

    for head != nil {
        newTail := ListNode{
            head.Val,
            tail,
        }
        tail = &newTail
        head = head.Next
    }

    return tail
}

