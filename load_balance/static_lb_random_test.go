package load_balance

import (
	"fmt"
	"testing"
	"time"
)

func TestRandomLB_GetServer(t *testing.T) {
	lb := NewRandomLoadBalancer(
		[]*RServer{
			{"10.13.0.5", 3},
			{"10.13.0.9", 1},
			{"10.13.0.7", 1},
		})

	for i := 0; i < 10; i++ {
		go func() {
			// s := lb.GetServer()
			s := lb.GetWeightedServer()
			fmt.Printf("%v\n", s.addr)
		}()
		time.Sleep(200 * time.Millisecond)
	}}