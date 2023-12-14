package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func parseCheck(line string, timeout int) int{
	failures := 0
	hostsAndPorts := strings.Split(line, " ")
	allHosts := hostsAndPorts[0]
	allPorts := hostsAndPorts[1]

	hosts := strings.Split(allHosts, ",")
	ports := strings.Split(allPorts, ",")
	for _, host := range hosts {
		for _, port := range ports {
			failures += checkHostAndPort(host, port, timeout)
		}

	}
	return failures
}

func checkHostAndPort(host string, port string, timeout int) int {
	failures := 0
	address := net.JoinHostPort(host, port)
	duration := time.Duration(int64(timeout) * int64(time.Second))
	con, err := net.DialTimeout("tcp", address, duration)
	if err != nil {
		fmt.Println("FAILED host:", host, " port:", port)
		failures++
	} else {
		if con != nil {
			fmt.Println("SUCCESS host:", host, " port:", port)
			con.Close()
		} else {
			fmt.Println("FAILED host:", host, " port:", port)
			failures++
		}

	}
	return failures
}

func main() {

	checkPtr := flag.String("check", "", "Formatted list of hosts and ports to check ex: h1,h2 80,9000")
	timeoutPtr := flag.Int("timeout", 3, "Timeout (in seconds); defaults to 3")
	flag.Parse()

	check := *checkPtr
	if check == "" {
		fmt.Println("Error: Must specify hosts and ports to check")
		os.Exit(125)
	}

	rc := parseCheck(*checkPtr, *timeoutPtr)
	os.Exit(rc)
}
