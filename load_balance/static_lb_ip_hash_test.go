package load_balance

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetHash(t *testing.T) {
	fmt.Println(GetHash("new"))
	fmt.Println(GetHash("new2"))
	fmt.Println(GetHash("new3") % 3)
}

func TestIPHashLB_GetServer(t *testing.T) {
	lb := &IPHashLB{
		servers: []*HServer{
			{"10.13.0.5"},
			{"10.13.0.9"},
			{"10.13.0.7"},
		}}
	N := 15
	for i := 0; i < N; i++ {
		iip := RandomIP()
		s := lb.GetServer(iip)
		fmt.Printf("%v -> %v\n", iip, s.addr)
		time.Sleep(200 * time.Millisecond)
	}
}

func RandomIP() string {
	rand.Seed(time.Now().UnixNano())
	ip := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
	return ip
}
