package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	Minimum int = iota
	Maximum
	Balanced
)

const file = "output.txt"

type Network struct {
	IP      string
	Subnets []Subnet
}

type Subnets []Subnet

type Subnet struct {
	Name      string
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
	network, err := promptNetwork()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	subnets, err := calculateSubnets(network)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := output(network.IP, subnets)
	err = ioutil.WriteFile(file, buf, 644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Subnets configuration saved to file: %s\n", file)
}
