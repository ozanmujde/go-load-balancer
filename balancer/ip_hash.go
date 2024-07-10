package balancer

import (
	"hash/fnv"
)

func ipHash(clientIp string, server []string) string {
	return server[hash(clientIp)%len(server)]
}

func hash(s string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return 0
	}
	return int(h.Sum32())
}
