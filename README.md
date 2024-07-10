# Go TCP Load Balancer

This is a simple TCP load balancer implemented in Go, supporting multiple load balancing strategies such as Round Robin, Random, Least Connections, and IP Hash. The load balancer distributes incoming connections across a set of backend servers.

## Features

- **Round Robin**: Distributes connections evenly across all servers in a circular order.
- **Random**: Selects a server at random for each incoming connection.
- **Least Connections**: Selects the server with the fewest active connections.
- **IP Hash**: Maps clients to servers based on the hash of their IP address, ensuring the same client consistently connects to the same server.


## Usage

### Setting Up

1. Clone the repository:
   ```bash
   $git clone <repository-url>
   $cd <repository-directory>
2. Changing the Load Balancing Strategy
You can change the load balancing strategy by calling the SetLoadBalancingStrategy function in main.go. The available strategies are:

    round_robin
    random
    least_connections
    ip_hash
Example:
   ```go
    balancer.SetLoadBalancingStrategy("least_connections")
3. Run
    ```bash
   $go run .
    or
    $go build -o load_balancer main.go
    $./load_balancer

   