package load_balance

import (
    "log"
    "sync"
)

type Server struct {
    Ip   string
    port int
}

type RoundRobin struct {
    sync.Mutex
    curr    int
    servers []*Server
}

func NewRoundRobin(servers []*Server) *RoundRobin {
    return &RoundRobin{
        curr:    0,
        servers: servers,
    }
}

func (lb *RoundRobin) GetServer() *Server {
    lb.Lock()
    defer lb.Unlock()
    log.Println(lb.curr)
    server := lb.servers[lb.curr]
    lb.curr = (lb.curr + 1) % len(lb.servers)
    return server
}
