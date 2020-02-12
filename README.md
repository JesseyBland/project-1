# Project-1

## Network Virtualization Functions


## Configuration

- Servers are added to the Config.yml file.
- First listed server in the Config.yml is the reverse proxy        Hostname and Port
- Under Reverse Proxy is listed servers that you add your back      end servers to.

## Reverse Proxy
- Reverse Proxy will default to the first listed Server under       servers: in the config.yml

## LoadBalanced Proxy
- Loadbalancer balances based of random number. The                 LoadBalancer will hit all availabe listed servers.

## noGO Firewall
- Blocks the incoming transmissions that are backend server         ports. Only allows client access to the proxy port.

## connTrace 
- Events on the proxy are traced on the conntrace tcp connection. 


## Features

- [x] Reverse Proxy
- [x] Load Balancer
- [x] Firewall
- [x] connTrace
