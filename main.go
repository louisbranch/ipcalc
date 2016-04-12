package main

import (
	"bufio"
	"fmt"
	"log"
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
	Size int
	Mode Mode
}

type Network struct {
	IPNet   *net.IPNet
	Subnets []Subnet
}

var in = bufio.NewScanner(os.Stdin)

func main() {
}

func prompt(msg string) string {
	fmt.Printf("%s: ", msg)
	if !in.Scan() {
		log.Fatalln(in.Err())
	}
	return in.Text()
}
