// Package linked_list
// Node.FindMidNode 寻找链表中点
// Node.RemoveLastKNode 删除链表倒数第K个节点，给定K合法
package linked_list

// FindMidNode 寻找链表中点
func (head *Node) FindMidNode() *Node {
    fast, slow := head, head
    // 如果链表长度为N，快指针如果移动了N，那么慢指针移动N/2
    // 中点就是最终慢指针的位置
    for fast != nil && fast.next != nil {
        fast = fast.next.next
        slow = slow.next
    }

    return slow
}

// RemoveLastKNode 删除链表倒数第K个节点，给定K合法
func (head *Node) RemoveLastKNode(K int) *Node {
    fast, slow := head, head

    // 快指针先跑K步
    for K > 0 {
        fast = fast.next
        K--
    }

    // 如果此时fast为nil，说明已经到达链表末尾了
    // 第一个节点就是倒数第n个几点
    if fast == nil {
        return head.next
    }

    // 同步前进
    for fast.next != nil {
        fast = fast.next
        slow = slow.next
    }

    // 删除节点
    slow.next = slow.next.next
    return head
}