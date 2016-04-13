package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
)

func calculateSubnets(network Network) ([]Subnet, error) {
	ip, ipnet, err := net.ParseCIDR(network.IP)
	if err != nil {
		return nil, err
	}

	var subnets []Subnet

	mask := ipnet.Mask
	ones, bits := mask.Size()
	max := bits - ones

	for _, subnet := range network.Subnets {
		size := subnet.Size + 2

		req := int(math.Ceil(math.Log2(float64(size))))
		size = pow2(req)

		if req > max {
			msg := "Network is too small. Subnet requested %d, max available is %d"
			return nil, fmt.Errorf(msg, subnet.Size, (pow2(max))-2)
		}

		n := binary.BigEndian.Uint32(ip.To4()) + uint32(size)
		ip = net.IPv4(byte(n>>24), byte(n>>16), byte(n>>8), byte(n))

		subnet.Size = size
		subnet.Mask = net.IP(mask).String()
		subnet.IP = ip.String()

		//fmt.Printf("IP: %s %d\n", ip, num)

		subnets = append(subnets, subnet)
	}

	return subnets, nil
}

func pow2(n int) int {
	return int(math.Pow(2, float64(n)))
}

func intToIP(n int) net.IP {
	ip := make(net.IP, 4)
	ip[0] = byte(n)
	ip[1] = byte(n >> 8)
	ip[2] = byte(n >> 16)
	ip[3] = byte(n >> 24)
	return ip
}
