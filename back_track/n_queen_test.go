package back_track

import (
    "fmt"
    "testing"
)

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