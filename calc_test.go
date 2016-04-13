package main

import (
	"reflect"
	"testing"
)

func TestCalculateSubnets(t *testing.T) {
	egs := []struct {
		network Network
		subnets []Subnet
		err     string
	}{
		{
			network: Network{IP: "11.22.33.44"},
			err:     "invalid CIDR address: 11.22.33.44",
		},
		{
			network: Network{
				IP: "11.22.33.44/24",
				Subnets: []Subnet{
					{
						Mode: Minimum,
						Size: 255,
					},
				},
			},
			err: "Network is too small. Subnet requested 255, max available is 254",
		},
		{
			network: Network{
				IP: "11.22.33.44/24",
				Subnets: []Subnet{
					{
						Mode: Minimum,
						Size: 10,
					},
				},
			},
			subnets: []Subnet{
				{
					Mode:      Minimum,
					Size:      16,
					IP:        "11.22.33.0/28",
					Mask:      "255.255.255.240",
					RangeMin:  "11.22.33.1",
					RangeMax:  "11.22.33.14",
					Broadcast: "11.22.33.15",
				},
			},
		},
	}

	for _, eg := range egs {
		got, err := calculateSubnets(eg.network)
		if err != nil && eg.err != err.Error() {
			t.Errorf("CalculateSubnets(%v)\n exp: error %s\n got: %s\n\n", eg.network, eg.err, err)
			continue
		}

		if !reflect.DeepEqual(eg.subnets, got) {
			t.Errorf("CalculateSubnets(%v)\n exp: %v\n got: %v\n\n", eg.network, eg.subnets, got)
		}

	}
}
