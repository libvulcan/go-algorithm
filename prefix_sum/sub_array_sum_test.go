package prefix_sum

import (
    "fmt"
    "testing"
)

func TestSubArraySum(t *testing.T) {
    fmt.Println(SubArraySum([]int{1, 1, 1}, 2))
}

func TestSubArraySumBetter(t *testing.T) {
    fmt.Println(SubArraySumBetter([]int{1, 1, 1}, 2))
}
