package sort

import (
    "fmt"
    "testing"
)

func TestRadixSort(t *testing.T) {
    arr := []int{5, 7, 2, 5, 6, 8, 4, 13, 5, 6, 7}

    RadixSort(&arr, MaxNumber(arr))
    fmt.Println(arr)
}