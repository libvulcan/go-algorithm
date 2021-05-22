// Package double_pointer
// RemoveDuplicates 去重
// DeleteDuplicates 链表去重
package double_pointer

type Node struct {
    val int
    next *Node
}
// RemoveDuplicates 数组去重
// nums 已排序数组
func RemoveDuplicates(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }

    slow, fast := 0, 0
    for fast < len(nums) {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
        fast++
    }
    return nums[:slow+1]
}

// DeleteDuplicates 链表去重
func DeleteDuplicates(head *Node) *Node {
    if head == nil || head.next == nil {
        return head
    }

    slow, fast := head, head
    for fast != nil {
        if fast.val != slow.val {
            slow.next = fast
            slow = slow.next
        }
        fast = fast.next
    }
    slow.next = nil
    return head
}
