package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Mode int

const (
	Minimum Mode = iota
	Maximum
	Balanced
)

type Subnet struct {
	Mode          Mode
	RequestedSize int
	RequiredSize  int
	IPNet         *net.IPNet
	HostMin       net.IP
	HostMax       net.IP
	Broadcast     net.IP
}

type Network struct {
	IPNet   *net.IPNet
	Subnets []Subnet
	Size    int
}

var in *bufio.Scanner

func init() {
	in = bufio.NewScanner(os.Stdin)
}

func main() {
	n, err := promptNetworkInfo()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(output(n))
}
