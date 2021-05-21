package back_track

import (
    "fmt"
    "testing"
)

func TestKNumberCombine(t *testing.T) {
    // 1~5，2个数的所有组合
    // K = 0 ([])
    // K = 1 ([1])                     ([2])               ([3])         ([4])
    // K = 2 ([1 2] [1 3] [1 4] [1 5]) ([2 3] [2 4] [2 5]) ([3 4] [3 5]) ([4 5])
	fmt.Println(KNumberCombine(5, 2))

	// 1~4，3个数的所有组合
	fmt.Println(KNumberCombine(4, 3))
}