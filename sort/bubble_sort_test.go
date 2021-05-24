package sort

import (
    "fmt"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    arr := []int{5, 7, 2, 5, 6, 8, 4, 13, 5, 6, 7}
    arr2 := append([]int{}, arr...)
    BubbleSort(&arr)
    fmt.Println(arr)

    BubbleSortWithChangeFlag(&arr2)
    fmt.Println(arr2)

}
