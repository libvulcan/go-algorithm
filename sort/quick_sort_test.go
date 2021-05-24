package sort

import (
    "fmt"
    "testing"
)

func TestQuickSort(t *testing.T) {
    arr := []int{5, 7, 2, 5, 6, 8, 4, 13, 5, 6, 7}

    QuickSort(&arr, 0, len(arr) - 1)
    fmt.Println(arr)
}
