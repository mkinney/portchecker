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
	//fmt.Println("hostsAndPorts:", hostsAndPorts)
	allHosts := hostsAndPorts[0]
	allPorts := hostsAndPorts[1]
	//fmt.Println("allHosts:", allHosts)
	//fmt.Println("allPorts:", allPorts)

	hosts := strings.Split(allHosts, ",")
	ports := strings.Split(allPorts, ",")
	for _, host := range hosts {
		//fmt.Println("host:", host)
		for _, port := range ports {
			failures += checkHostAndPort(host, port, timeout)
		}

	}
	return failures
}

func checkHostAndPort(host string, port string, timeout int) int {
	failures := 0
	//fmt.Println("check host:", host, " port:", port)
	address := net.JoinHostPort(host, port)
	duration := time.Duration(int64(timeout) * int64(time.Second))
	//fmt.Println("duration:", duration)
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

	//fmt.Println("check:", *checkPtr)
	//fmt.Println("timeout:", *timeoutPtr)

	rc := parseCheck(*checkPtr, *timeoutPtr)
	os.Exit(rc)
}
