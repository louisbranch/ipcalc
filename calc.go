package main

import (
	"fmt"
	"math"
	"net"
)

func calculateSubnets(network Network) ([]Subnet, error) {
	_, ipnet, err := net.ParseCIDR(network.IP)
	if err != nil {
		return nil, err
	}

	var subnets []Subnet

	mask := ipnet.Mask

	for _, subnet := range network.Subnets {
		size := subnet.Size + 2
		ones, bits := mask.Size()
		max := bits - ones

		req := int(math.Ceil(math.Log2(float64(size))))
		size = pow2(req)

		if req > max {
			msg := "Network is too small. Subnet requested %d, max available is %d"
			return nil, fmt.Errorf(msg, subnet.Size, (pow2(max))-2)
		}

		subnet.Size = size

		fmt.Printf("Required %d, ones: %d, bits: %d\n", req, ones, bits)

		subnets = append(subnets, subnet)
	}

	return subnets, nil
}

func pow2(n int) int {
	return int(math.Pow(2, float64(n)))
}
