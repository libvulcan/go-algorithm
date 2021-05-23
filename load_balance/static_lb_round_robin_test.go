package load_balance

import (
    "fmt"
    "testing"
    "time"
)

func TestRoundRobin(t *testing.T) {
    lb := NewRoundRobin([]*Server{
        {"10.11.0.1", 2233}, {"10.11.0.5", 2233}, {"10.11.0.9", 2235},
    })

    for i := 0; i < 10; i++ {
        go func() {
            s := lb.GetServer()
            fmt.Printf("%v:%v", s.Ip, s.port)
        }()
        time.Sleep(200 * time.Millisecond)
    }
}

func BenchmarkRoundRobin(b *testing.B) {
    lb := NewRoundRobin([]*Server{
        {"10.11.0.1", 2233}, {"10.11.0.5", 2233}, {"10.11.0.9", 2235},
    })
    for i := 0; i < b.N; i++ {
        go func() {
            lb.GetServer()
            // fmt.Printf("%v:%v", s.Ip, s.port)
        }()
    }
}
