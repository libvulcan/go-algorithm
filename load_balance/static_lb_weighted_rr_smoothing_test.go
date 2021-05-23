package load_balance

import (
    "fmt"
    "testing"
    "time"
)

func TestSmoothWeightedRoundRobin(t *testing.T) {
    lb := NewSmoothWeightedRoundRobin([]*SmoothWServer{
        {"10.11.0.1", 3, 0}, {"10.11.0.5", 1, 0}, {"10.11.0.3", 1, 0},
    })

    for i := 0; i < 5; i++ {
        // go func() {
        s := lb.GetServer()
        fmt.Printf("%v:%v\n", s.addr, s.weight)
        // }()
        time.Sleep(200 * time.Millisecond)
    }
}