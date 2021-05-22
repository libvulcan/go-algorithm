package binary_search

import (
    "fmt"
    "testing"
)

var arr = []int{1, 3, 3, 5, 7, 8, 12, 13}
var arr2 = []int{1, 2, 2, 2, 3}

func TestFindIndex(t *testing.T) {
    // 返回5
    fmt.Println(FindIndex(arr, 8))

    // 返回2，实际上返回1、2、3都对应target=2
    // 无法处理边界问题
    fmt.Println(FindIndex(arr2, 2))
}

func TestFindLeftBoundIndexFirst(t *testing.T) {
    // 返回左边界第一个
    fmt.Println(FindLeftBoundIndexFirst(arr2, 2))
}

func TestFindRightBoundIndexFirst(t *testing.T) {
    fmt.Println(FindRightBoundIndexFirst(arr2, 2))
}