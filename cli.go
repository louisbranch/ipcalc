package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
	"text/tabwriter"
)

func promptNetworkInfo() (network Network, err error) {
	input := prompt("IPv4 Network address (CIDR format)")

	_, ipnet, err := net.ParseCIDR(input)
	if err != nil {
		return network, fmt.Errorf("Error parsing network address %s", err)
	}

	network.IPNet = ipnet

	input = prompt("Number of subnets to create?")
	n, err := strconv.Atoi(input)
	if err != nil {
		return network, fmt.Errorf("Error parsing number of subnets %s", err)
	}

	for i := 0; i < n; i++ {
		subnet, err := promptSubnet(i)
		if err != nil {
			return network, err
		}
		network.Subnets = append(network.Subnets, subnet)
	}

	return network, nil
}

func promptSubnet(i int) (subnet Subnet, err error) {
	msg := fmt.Sprintf("Subnet #%d size", i+1)
	input := prompt(msg)
	n, err := strconv.Atoi(input)
	if err != nil {
		return subnet, fmt.Errorf("Error parsing subnet size %s", err)
	}
	if n < 1 {
		return subnet, fmt.Errorf("Subnet size must be greater than 0")
	}
	subnet.RequestedSize = n

	msg = fmt.Sprintf("Subnet #%d mode (0 = Mininum, 1 = Maximum, 2 = Balanced)", i+1)
	input = prompt(msg)
	n, err = strconv.Atoi(input)
	if err != nil {
		return subnet, fmt.Errorf("Error parsing subnet mode %s", err)
	}
	if n < 0 || n > 3 {
		return subnet, fmt.Errorf("Invalid mode")
	}
	subnet.Mode = Mode(n)

	return subnet, nil
}

func output(res Network) string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 8, 8, 2, ' ', 0)
	fmt.Fprintf(w, "Address:\t%s\n", res.IPNet)

	mask := res.IPNet.Mask
	fmt.Fprintf(w, "Netmask:\t%d.%d.%d.%d\n", mask[0], mask[1], mask[2], mask[3])

	w.Flush()
	return buf.String()
}

func prompt(msg string) string {
	fmt.Printf("%s: ", msg)
	if !in.Scan() {
		log.Fatalln(in.Err())
	}
	return in.Text()
}
