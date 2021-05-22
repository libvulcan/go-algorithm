package linked_list

import (
    "fmt"
    "testing"
)

var lk = &Node{
    val: 1,
    next: &Node{
        val: 2,
        next: &Node{
            val: 3,
            next: &Node{
                val: 4,
                next: &Node{
                    val:  5,
                    next: nil,
                },
            },
        },
    },
}

func TestFindMidNode(t *testing.T) {
    fmt.Println(lk.FindMidNode().val)
}

func TestRemoveLastKNode(t *testing.T) {
    // 链表1 -> 2 -> 3 -> 4 -> 5 -> nil
    // 移除倒数第3个数，即移除&Node{val: 3}
    // 输出 1 -> 2 -> 4 -> 5 -> nil
    res := lk.RemoveLastKNode(3)
    for res != nil {
        fmt.Print(res.val, " -> ")
        res = res.next
    }

    fmt.Println("nil")
}
