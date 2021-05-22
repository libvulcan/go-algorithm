package back_track

import (
    "fmt"
    "testing"
)

func TestCanPartitionKSubSet(t *testing.T) {
    fmt.Println(CanPartitionKSubSet([]int{1, 1, 2, 2, 3, 3, 1, 2, 3}, 6))
}

func TestCanPartitionKSubSetFromBucket(t *testing.T) {
    fmt.Println(CanPartitionKSubSetFromBucket([]int{1, 1, 2, 2, 3, 3, 1, 2, 3}, 6))
}

func TestAllSubsets(t *testing.T) {
    fmt.Println(AllSubsets([]int{1, 2, 3, 4}))
}
