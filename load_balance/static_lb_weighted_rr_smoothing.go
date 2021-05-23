package load_balance

import "sync"

type SmoothWServer struct {
    addr       string
    weight     int64
    currWeight int64
}

type SmoothWRoundRobin struct {
    sync.Mutex
    servers    []*SmoothWServer
}

func NewSmoothWeightedRoundRobin(servers []*SmoothWServer) *SmoothWRoundRobin {
    return &SmoothWRoundRobin{
        servers: servers,
    }
}

// 平滑轮询
func (lb *SmoothWRoundRobin) GetServer() *SmoothWServer {
    lb.Lock()
    defer lb.Unlock()

    var total int64
    var index int = -1
    for i, v := range lb.servers {
        v.currWeight += v.weight
        total += v.weight

        if index == -1 || lb.servers[index].currWeight < v.currWeight {
            index = i
        }
    }

    lb.servers[index].currWeight -= total
    return lb.servers[index]
}
