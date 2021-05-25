package sort

import (
    "fmt"
    "testing"
)

func TestHeapSort(t *testing.T) {
    heap := &Heap{data: []int{5, 7, 2, 5, 6, 8, 4, 13, 5, 6, 7}}
    heap.Sort()
    fmt.Println(heap.data)
}
