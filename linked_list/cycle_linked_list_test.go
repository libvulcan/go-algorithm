package linked_list

import (
    "fmt"
    "testing"
)

func TestCycleLinkedList(t *testing.T) {

    node5 := &Node{
        val:  5,
        next: nil,
    }

    node3 := &Node{
        val: 3,
        next: &Node{
            val:  4,
            next: nil,
        },
    }

    // 连成链表，node4.next = node5
    node3.next.next = node5
    // 构造环，环的开始节点是node3
    node5.next = node3

    lk := &Node{
        val: 1,
        next: &Node{
            val:  2,
            next: node3,
        },
    }

    fmt.Println(lk.HasCycle())
    fmt.Println(lk.GetCycle().val)
}
