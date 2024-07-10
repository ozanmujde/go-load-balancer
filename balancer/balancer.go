package balancer

import "sync"

var (
	strategy = "round_robin"
)

func ChooseServer(server []string, serverMetrics map[string]int, clientIp string, mu *sync.Mutex) string {
	mu.Lock()
	defer mu.Unlock()
	switch strategy {
	case "round_robin":
		return roundRobin(server)
	case "random":
		return randomChoice(server)
	case "least_connections":
		return leastConnections(server, serverMetrics)
	case "ip_hash":
		return ipHash(clientIp, server)
	default:
		return roundRobin(server)
	}
}

func GetLoadBalancingStrategy() string {
	return strategy
}

func SetLoadBalancingStrategy(balancingStrategy string) string {
	strategy = balancingStrategy
	return strategy
}
