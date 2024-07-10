package balancer

var count int

func roundRobin(server []string) string {
	s := server[count%len(server)]
	count++
	return s
}
