package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
	"text/tabwriter"
)

func promptNetworkInfo() (Network, error) {
	req := Network{}

	addr := prompt("IPv4 Network address (CIDR format)")
	nets := prompt("How many subnets you want to create?")

	_, ipnet, err := net.ParseCIDR(addr)
	if err != nil {
		return req, fmt.Errorf("Error parsing network address %s", err)
	}

	req.IPNet = ipnet

	size, err := strconv.Atoi(nets)
	if err != nil {
		return req, err
	}

	for i := 0; i < size; i++ {
		n := prompt(fmt.Sprintf("Subnet #%d size", i+1))
		_, err := strconv.Atoi(n)
		if err != nil {
			log.Fatalln(err)
		}
		n = prompt(fmt.Sprintf("Subnet #%d mode (0 = Mininum, 1 = Maximum, 2 = Balanced)", i+1))
		_, err = strconv.Atoi(n)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return Network{}, nil
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
