package back_track

import (
	"fmt"
	"testing"
)

func TestFullPer(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	var track []int
	var res [][]int
	FullPermutation(&res, nums, track)
	fmt.Println(len(res))
}

func TestNQueen(t *testing.T) {
	N := 8
	var res [][][]int
	var track [][]int

	for i := 0; i < N; i++ {
		track = append(track, make([]int, N))
	}

	NQueen(&res, track, N-1)

	fmt.Println(len(res))
}

func TestCanPartitionKSubSet(t *testing.T) {
	fmt.Println(CanPartitionKSubSet([]int{1, 1, 2, 2, 3, 3, 1, 2, 3}, 6))
}

func TestCanPartitionKSubSetFromBucket(t *testing.T) {
	fmt.Println(CanPartitionKSubSetFromBucket([]int{1, 1, 2, 2, 3, 3, 1, 2, 3}, 6))
}
