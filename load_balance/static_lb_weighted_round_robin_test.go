package load_balance

import (
    "fmt"
    "testing"
    "time"
)

func TestWeightedRoundRobin(t *testing.T) {
    lb := NewWeightedRoundRobin([]*WeightedServer{
        {"10.11.0.1", 3}, {"10.11.0.5", 1}, {"10.11.0.3", 1},
    })

    for i := 0; i < 5; i++ {
        // go func() {
        s := lb.GetServer()
        fmt.Printf("%v:%v\n", s.addr, s.weight)
        // }()
        time.Sleep(200 * time.Millisecond)
    }
}