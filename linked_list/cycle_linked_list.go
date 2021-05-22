// Package linked_list
// Node.HasCycle 判断链表中是否有环
// Node.GetCycle 获取环的初始位置
package linked_list

// HasCycle 判断链表中是否有环
// 快慢指针遍历
func (head *Node) HasCycle() bool {
    fast, slow := head, head
    for fast != nil && fast.next != nil {
        fast = fast.next.next
        slow = slow.next
        // 如果有环，快指针最终会追上慢指针
        if fast == slow {
            return true
        }
    }
    return false
}

// GetCycle 获取环的初始位置
func (head *Node) GetCycle() *Node {
    fast, slow := head, head
    for fast != nil && fast.next != nil {
        fast = fast.next.next
        slow = slow.next
        // 第一次相遇时跳出遍历
        if fast == slow {
            break
        }
    }

    // 此时若slow移动了k步，那么fast移动了2k步
    // 假设第一次相遇后所在节点距离head节点为m
    // 那么head到起始节点的距离是k-m
    // 将fast指向head，fast走k-m步就能到达环的起始节点
    fast = head
    // 以同样的速度继续遍历
    // k-m步后会fast和slow会再次相遇
    // 此时的位置就是环的起始位置
    for slow != fast {
        fast = fast.next
        // slow从相遇点开始走，同理走k-m步也到环起始节点
        slow = slow.next
    }
    return fast
}