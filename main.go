package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	// dialTimeout is the timeout for dialing a host:port.
	dialTimeout = time.Duration(3) * time.Second

	// sleepTime is the time to sleep between dial attempts.
	sleepTime = func() { time.Sleep(3 * time.Second) }
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "Error: you need to provide a host and port to test.")
		os.Exit(1)
	}

	host := os.Args[1] // host or ip
	port := os.Args[2]

	durationValue, err := strconv.Atoi(os.Args[3]) // duration in seconds to keep trying
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid timeout value.")
		os.Exit(1)
	}

	duration := time.Duration(durationValue) * time.Second
	addr := net.JoinHostPort(host, port)

	if await(addr, duration) {
		fmt.Printf("Successfully connected to '%s'.\n", addr)
		os.Exit(0)
	}

	fmt.Printf("Failed to connect to '%s'.\n", addr)
	os.Exit(1)
}

func await(addr string, duration time.Duration) bool {
	startTime := time.Now()

	for time.Since(startTime) < duration {
		conn, err := net.DialTimeout("tcp", addr, dialTimeout)
		if err == nil {
			conn.Close()
			return true
		}
		sleepTime()
	}
	return false
}
