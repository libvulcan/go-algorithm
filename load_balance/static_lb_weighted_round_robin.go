package load_balance

import (
    "go-algorithm/load_balance/math"
    "sync"
)

type WeightedServer struct {
    addr   string
    weight int64
}

type WeightedRoundRobin struct {
    sync.Mutex
    currPos    int
    currWeight int64
    servers    []*WeightedServer
}

func NewWeightedRoundRobin(servers []*WeightedServer) *WeightedRoundRobin {
    return &WeightedRoundRobin{
        currPos:    -1,
        currWeight: 0,
        servers:    servers,
    }
}

func MaxWeight(servers []*WeightedServer) int64 {
    var max int64 = 0
    for _, v := range servers {
        if v.weight >= max {
            max = v.weight
        }
    }
    return max
}

// GetServerSmoothing 平滑方案：解决连续选取问题
func (lb *WeightedRoundRobin) GetServerSmoothing() *WeightedServer {
    return nil
}

// GetServer 普通方案
func (lb *WeightedRoundRobin) GetServer() *WeightedServer {
    lb.Lock()
    defer lb.Unlock()
    gcd := lb.GCDWeight(len(lb.servers) - 1)
    for {
        lb.currPos = (lb.currPos + 1) % len(lb.servers)
        if lb.currPos == 0 {
            lb.currWeight -= gcd
            if lb.currWeight <= 0 {
                lb.currWeight = MaxWeight(lb.servers)
                if lb.currWeight == 0 {
                    return nil
                }
            }
        }

        if lb.servers[lb.currPos].weight >= lb.currWeight {
            return lb.servers[lb.currPos]
        }
    }
}

func (lb *WeightedRoundRobin) GCDWeight(n int) int64 {
    if n == 0 {
        return lb.servers[n].weight
    } else {
        return math.GCD(lb.servers[n].weight, lb.GCDWeight(n-1))
    }
}
