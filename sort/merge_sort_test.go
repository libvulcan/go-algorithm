package sort

import (
    "fmt"
    "testing"
)

func TestMergeSort(t *testing.T) {
    arr := []int{5, 7, 2, 5, 6, 8, 4, 13, 5, 6, 7}

    // 子序列 = len(当前序列) / 2
    // {5, 7, 2, 5, 6}      {8, 4, 13, 5, 6, 7}
    // {5, 7}   {2, 5, 6}   {8, 4, 13} {5, 6, 7}
    // {5} {7}  {2} {5, 6}  {8} {4, 13} {5} {6, 7}
    // 归并: 调用栈
    // l = func(func(5,7) func(2, func(5,6)))
    // r = func(func(8, func(4,13)) func(5, fun(6,7)))
    // fun(l, r)
    fmt.Println(MergeSort(arr))
}