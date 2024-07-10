package balancer

import (
	"errors"
	"math"
	"sync"
)

func leastConnections(server []string, serverMetrics map[string]int) string {
	var leastConnectedServer string
	leastConnection := math.MaxInt
	for _, srv := range server {
		if serverMetrics[srv] < leastConnection {
			leastConnection = serverMetrics[srv]
			leastConnectedServer = srv
		}
	}
	serverMetrics[leastConnectedServer]++
	return leastConnectedServer
}

func DecreaseConnections(server string, serverMetrics map[string]int, mu *sync.Mutex) error {
	mu.Lock()
	defer mu.Unlock()
	serverMetrics[server]--
	if serverMetrics[server] < 0 {
		serverMetrics[server] = 0
		return errors.New("server connection count went below zero")
	}
	return nil
}
