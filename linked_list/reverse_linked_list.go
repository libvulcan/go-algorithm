// Package linked_list
// ReverseLinkedList 反转整个链表
// ReverseLinkedListN 反转链表前N个节点
// ReverseLinkedListBetween 反转链表N~M节点
// Reverse 反转链表，双指针迭代
package linked_list

type Node struct {
	val  int
	next *Node
}

// ReverseLinkedList 反转整个链表
func ReverseLinkedList(head *Node) *Node {
	if head.next == nil {
		return head
	}

	last := ReverseLinkedList(head.next)
	head.next.next = head
	head.next = nil
	return last
}

// ReverseLinkedListN 反转链表前N个节点
var tail *Node

func ReverseLinkedListN(head *Node, N int) *Node {
	if N == 1 {
		tail = head.next
		return head
	}

	last := ReverseLinkedListN(head.next, N-1)
	head.next.next = head
	head.next = tail
	return last
}

// ReverseLinkedListBetween 反转链表N~M节点
func ReverseLinkedListBetween(head *Node, M, N int) *Node {
	if M == 1 {
		return ReverseLinkedListN(head, N)
	}

	head.next = ReverseLinkedListBetween(head.next, M-1, N-1)
	return head
}

// Reverse 反转链表，双指针迭代
func Reverse(head *Node, M int, N int) *Node {
	// 如果只有一个节点或者头节点为空，则直接返回head指针
	// 或者只翻转第一个节点
	if head == nil || head.next == nil {
		return head
	}

	// 声明一个指针充当新的头节点
	h := &Node{}
	h.next = head

	// 当前指针指向：头指针的下一个节点
	curr := head
	// 结尾指针指向：当前指针的下一个节点
	end := curr.next

	// 当前迭代次数
	var currIndex int = 1

	// 开始迭代位置的上一个节点指针
	var last *Node

	// 迭代
	for end != nil {

		if currIndex < M && M != -1 {
			last = head
			head, curr = head.next, head.next
			end = curr.next
			currIndex++
			continue
		}

		if currIndex > N - 1 && N != -1 {
			break
		}

		// 翻转curr和end指针指向节点的位置
		curr.next, end.next = end.next, head
		head, end = end, curr.next

		if last != nil {
			last.next = head
		}

		currIndex++
	}

	// 如果从头开始翻转，翻转完后head指向的仍然是表头
	if M == 1 || M == -1{
		return head
	}

	return h.next
}
