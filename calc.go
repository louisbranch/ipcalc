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
	wildMask := bits - ones
	ip = ip.Mask(mask)

	for _, subnet := range network.Subnets {
		size := subnet.Size + 2 //host && broadcast

		reqBits := int(math.Ceil(math.Log2(float64(size))))
		size = pow2(reqBits)

		if reqBits > wildMask {
			msg := "Network is too small. Subnet requested %d, max available is %d"
			return nil, fmt.Errorf(msg, subnet.Size, (pow2(wildMask))-2)
		}

		mask = net.CIDRMask(bits-reqBits, bits)
		ones, _ = mask.Size()

		subnet.Size = size
		subnet.IP = fmt.Sprintf("%s/%d", ip, ones)
		subnet.Mask = net.IP(mask).String()

		n := binary.BigEndian.Uint32(ip.To4())

		n += 1 // host
		subnet.RangeMin = intToIP(n).String()

		n += uint32(size) - 1 // last bit
		ip = intToIP(n)       // new ip

		n -= 1
		subnet.Broadcast = intToIP(n).String()

		n -= 1
		subnet.RangeMax = intToIP(n).String()

		subnets = append(subnets, subnet)
	}

	return subnets, nil
}

func pow2(n int) int {
	return int(math.Pow(2, float64(n)))
}

func intToIP(n uint32) net.IP {
	ip := make(net.IP, 4)
	ip[0] = byte(n >> 24)
	ip[1] = byte(n >> 16)
	ip[2] = byte(n >> 8)
	ip[3] = byte(n)
	return ip
}
