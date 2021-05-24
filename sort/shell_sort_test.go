package sort

import (
    "fmt"
    "testing"
)

func TestShellSort(t *testing.T) {
    arr := []int{5, 7, 2, 5, 6, 8, 4, 13, 5, 6, 7}

    ShellSort(&arr, 2)
    fmt.Println(arr)
}