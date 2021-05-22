package back_track

import (
    "fmt"
    "testing"
)

func TestFullPermutation(t *testing.T) {
    nums := []int{1, 2, 3, 4, 5, 6, 7}
    var track []int
    var res [][]int
    FullPermutation(&res, nums, track)
    fmt.Println(len(res))
}
