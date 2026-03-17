package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var nServers, nIPs int
	if _, err := fmt.Fscan(in, &nServers, &nIPs); err != nil {
		return
	}

	servers := make(map[string]string, nServers)
	for i := 0; i < nServers; i++ {
		var serverName, ip string
		if _, err := fmt.Fscan(in, &serverName, &ip); err != nil {
			return
		}
		servers[ip] = serverName
	}

	for i := 0; i < nIPs; i++ {
		var ipName, ipWithSemi string
		if _, err := fmt.Fscan(in, &ipName, &ipWithSemi); err != nil {
			return
		}
		ip := strings.TrimSuffix(ipWithSemi, ";")
		if server, ok := servers[ip]; ok {
			fmt.Printf("%s %s; #%s\n", ipName, ip, server)
		}
	}
}
