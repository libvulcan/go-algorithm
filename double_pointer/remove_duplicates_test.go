package double_pointer

import (
    "fmt"
    "testing"
)

func TestRemoveDuplicates(t *testing.T) {
    fmt.Println(RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 3, 4, 4, 5}))
    fmt.Println(RemoveDuplicates([]int{0, 1, 1, 1, 1, 1, 5, 5, 11, 23, 23, 23, 44}))
}

func TestDeleteDuplicates(t *testing.T) {
    head := &Node{
        val:  0,
        next: &Node{
            val:  1,
            next: &Node{
                val:  1,
                next: &Node{
                    val:  3,
                    next: &Node{
                        val:  4,
                        next: &Node{
                            val:  4,
                            next: nil,
                        },
                    },
                },
            },
        },
    }
    h := DeleteDuplicates(head)
    for h != nil {
        fmt.Print(h.val, " >> ")
        h = h.next
    }

    fmt.Println("nil")
}