# Project-1

## Network Virtualization Functions


## Requirements
Go is a a requirement and is available for download at https://golang.org

Windows users will need to use bash to build this project. Bash is incorporated into git and is available at https://gitforwindows.org/

## Installation
In an open terminal / bash, download the project with command: go get -u github.com/JesseyBland/project-1

## Configuration

- Servers are added to the Config.yml file.
- First listed server in the Config.yml is the reverse proxy        Hostname and Port
- Second listed server in the Config.yml is the load balancer       Hostname and Port
- Third listed server in the Config.yml is the connection tracer    Hostname and Port
- Under servers: listed servers that you add your backend servers to.

## Reverse Proxy
- Reverse Proxy will default to the first listed Server under servers: in the config.yml
- Default listens on port 6060

## LoadBalanced Proxy
- Loadbalancer balances based of random number. The LoadBalancer will hit all availabe listed servers at random.
- Default listens on port 6061
## noGO Firewall
- Blocks the incoming transmissions that are backend server ports. Only allows client access to the proxy port.

## connTrace 
- Events on the proxy are traced on the conntrace tcp connection. 
- Default listens on port 3333

## Features

- [x] Reverse Proxy
- [x] Load Balancer
- [x] Firewall
- [x] connTrace
