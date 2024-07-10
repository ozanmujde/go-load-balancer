package balancer

import "math/rand"

func randomChoice(server []string) string {
	return server[rand.Int()*len(server)]
}
