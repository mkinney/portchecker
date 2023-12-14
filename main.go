package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	rc := 0
	checkPtr := flag.String("check", "", "Formatted list of hosts and ports to check ex: h1,h2 80,9000")
	timeoutPtr := flag.Int("timeout", 3, "Timeout (in seconds); defaults to 3")
	flag.Parse()

	check := *checkPtr
	if check == "" {
		fmt.Println("Error: Must specify hosts and ports to check")
		os.Exit(125)
	}

	fmt.Println("check:", *checkPtr)
	fmt.Println("timeout:", *timeoutPtr)

	os.Exit(rc)
}
