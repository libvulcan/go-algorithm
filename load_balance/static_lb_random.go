package load_balance

import (
	"math/rand"
	"sync"
	"time"
)

type RServer struct {
	addr   string
	weight int64
}

type RandomLB struct {
	sync.Mutex
	servers     []*RServer
	totalWeight int64
}

func NewRandomLoadBalancer(servers []*RServer) *RandomLB {
	var totalWeight int64
	for _, v := range servers {
		totalWeight += v.weight
	}
	return &RandomLB{
		servers:     servers,
		totalWeight: totalWeight,
	}
}

func (lb *RandomLB) GetServer() *RServer {
	lb.Lock()
	defer lb.Unlock()

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(lb.servers))
	return lb.servers[index]
}

func (lb *RandomLB) GetWeightedServer() *RServer {
	lb.Lock()
	defer lb.Unlock()

	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Int63n(lb.totalWeight)

	var index int64 = 0
	for _, v := range lb.servers {
		index += v.weight
		if v.weight >= randIndex {
			return v
		}
	}

	return lb.servers[len(lb.servers)-1]
}
