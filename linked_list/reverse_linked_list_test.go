package linked_list

import (
	"fmt"
	"testing"
)

var linkedList = &Node{
	val: 0,
	next: &Node{
		val: 1,
		next: &Node{
			val: 2,
			next: &Node{
				val:  3,
				next: &Node{
					val:  4,
					next: nil,
				},
			},
		},
	},
}

func TestReverseLinkedList(t *testing.T) {
	reverseList := ReverseLinkedList(linkedList)
	for reverseList != nil {
		fmt.Println(reverseList.val)
		reverseList = reverseList.next
	}
}

func TestReverseLinkedListN(t *testing.T) {
	reverseList := ReverseLinkedListN(linkedList, 3)
	for reverseList != nil {
		fmt.Println(reverseList.val)
		reverseList = reverseList.next
	}
}

func TestReverseLinkedListBetween(t *testing.T) {
	reverseList := ReverseLinkedListBetween(linkedList, 1, 4)
	for reverseList != nil {
		fmt.Println(reverseList.val)
		reverseList = reverseList.next
	}
}

func TestReverse(t *testing.T) {
	// M=-1 表示从头开始
	// N=-1 表示一直翻转到结尾
	// 范围M~N
	reverseList := Reverse(linkedList, 2, -1)
	for reverseList != nil {
		fmt.Println(reverseList.val)
		reverseList = reverseList.next
	}
}

