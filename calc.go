package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
	"sort"
)

func calculateSubnets(network Network) (Subnets, error) {
	ip, ipnet, err := net.ParseCIDR(network.IP)
	if err != nil {
		return nil, err
	}

	subnets := validateModes(network.Subnets)

	mask := ipnet.Mask
	ones, bits := mask.Size()
	wildMask := bits - ones
	ip = ip.Mask(mask)

	for i, subnet := range subnets {
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

		subnets[i] = subnet
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

func (s Subnets) Less(i, j int) bool {
	if s[i].Mode == s[j].Mode {
		return s[i].Size < s[j].Size
	}
	return s[i].Mode < s[j].Mode
}

func (s Subnets) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Subnets) Len() int {
	return len(s)
}

func validateModes(subnets Subnets) Subnets {
	max := 0
	for _, subnet := range subnets {
		if subnet.Mode == Maximum {
			max += 1
		}
		if max > 1 {
			break
		}
	}

	if max > 1 {

		for i, subnet := range subnets {
			if subnet.Mode == Maximum {
				subnet.Mode = Balanced
			}
			subnets[i] = subnet
		}
	}

	sort.Sort(subnets)

	return subnets
}
