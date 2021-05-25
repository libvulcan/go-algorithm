package sort

import (
    "fmt"
    "testing"
)

func TestBucketSort(t *testing.T) {
    arr := []int{5, 7, 2, 5, 6, 8, 4, 13, 5, 6, 7}

    BucketSort(&arr, 5)
    fmt.Println(arr)
}