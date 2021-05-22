package double_pointer

import (
    "fmt"
    "testing"
)

func TestRemoveElement(t *testing.T) {
    fmt.Println(RemoveElement([]int{1, 1, 2, 3, 3, 4, 5, 6}, 3))
}
