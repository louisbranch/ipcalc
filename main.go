package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Minimum int = iota
	Maximum
	Balanced
)

type Network struct {
	IP      string
	Subnets []Subnet
}

type Subnet struct {
	Mode      int
	Size      int
	IP        string
	Mask      string
	RangeMin  string
	RangeMax  string
	Broadcast string
}

var in *bufio.Scanner

func init() {
	in = bufio.NewScanner(os.Stdin)
}

func main() {
	_, err := promptNetwork()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Println(output(n))
}
