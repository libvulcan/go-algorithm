package load_balance

import (
	"hash/crc32"
	"sync"
)

type HServer struct {
	addr string
}

type IPHashLB struct {
	sync.Mutex
	servers []*HServer
}

func (lb *IPHashLB) GetServer(ip string) *HServer {
	lb.Lock()
	defer lb.Unlock()

	index := GetHash(ip) % int32(len(lb.servers))
	return lb.servers[index]
}

func GetHash(str string) int32 {
	hashcode := int32(crc32.ChecksumIEEE([]byte(str)))
	if hashcode < 0 {
		hashcode = -hashcode
	}
	return hashcode
}