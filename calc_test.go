package main

import (
	"reflect"
	"testing"
)

func TestCalculateSubnets(t *testing.T) {
	egs := []struct {
		network Network
		subnets Subnets
		err     string
	}{
		{
			network: Network{IP: "11.22.33.44"},
			err:     "invalid CIDR address: 11.22.33.44",
		},
		{
			network: Network{
				IP: "11.22.33.44/24",
				Subnets: Subnets{
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
				Subnets: Subnets{
					{
						Mode: Minimum,
						Size: 10,
					},
				},
			},
			subnets: Subnets{
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
		{
			network: Network{
				IP: "192.168.1.2/30",
				Subnets: Subnets{
					{
						Mode: Minimum,
						Size: 2,
					},
				},
			},
			subnets: Subnets{
				{
					Mode:      Minimum,
					Size:      4,
					IP:        "192.168.1.0/30",
					Mask:      "255.255.255.252",
					RangeMin:  "192.168.1.1",
					RangeMax:  "192.168.1.2",
					Broadcast: "192.168.1.3",
				},
			},
		},
		{
			network: Network{
				IP: "192.137.28.3/26",
				Subnets: Subnets{
					{
						Mode: Minimum,
						Size: 20,
					},
				},
			},
			subnets: Subnets{
				{
					Mode:      Minimum,
					Size:      32,
					IP:        "192.137.28.0/27",
					Mask:      "255.255.255.224",
					RangeMin:  "192.137.28.1",
					RangeMax:  "192.137.28.30",
					Broadcast: "192.137.28.31",
				},
			},
		},
		{
			network: Network{
				IP: "192.137.28.3/18",
				Subnets: Subnets{
					{
						Mode: Minimum,
						Size: 300,
					},
				},
			},
			subnets: Subnets{
				{
					Mode:      Minimum,
					Size:      512,
					IP:        "192.137.0.0/23",
					Mask:      "255.255.254.0",
					RangeMin:  "192.137.0.1",
					RangeMax:  "192.137.1.254",
					Broadcast: "192.137.1.255",
				},
			},
		},
		{
			network: Network{
				IP: "11.22.33.44/24",
				Subnets: Subnets{
					{
						Mode: Minimum,
						Size: 10,
					},
					{
						Mode: Minimum,
						Size: 10,
					},
				},
			},
			subnets: Subnets{
				{
					Mode:      Minimum,
					Size:      16,
					IP:        "11.22.33.0/28",
					Mask:      "255.255.255.240",
					RangeMin:  "11.22.33.1",
					RangeMax:  "11.22.33.14",
					Broadcast: "11.22.33.15",
				},
				{
					Mode:      Minimum,
					Size:      16,
					IP:        "11.22.33.16/28",
					Mask:      "255.255.255.240",
					RangeMin:  "11.22.33.17",
					RangeMax:  "11.22.33.30",
					Broadcast: "11.22.33.31",
				},
			},
		},
		{
			network: Network{
				IP: "192.137.28.3/18",
				Subnets: Subnets{
					{
						Name: "IT",
						Mode: Minimum,
						Size: 300,
					},
					{
						Name: "Marketing",
						Mode: Minimum,
						Size: 30,
					},
					{
						Name: "Finance",
						Mode: Minimum,
						Size: 15,
					},
				},
			},
			subnets: Subnets{
				{
					Name:      "Finance",
					Mode:      Minimum,
					Size:      32,
					IP:        "192.137.0.0/27",
					Mask:      "255.255.255.224",
					RangeMin:  "192.137.0.1",
					RangeMax:  "192.137.0.30",
					Broadcast: "192.137.0.31",
				},
				{
					Name:      "Marketing",
					Mode:      Minimum,
					Size:      32,
					IP:        "192.137.0.32/27",
					Mask:      "255.255.255.224",
					RangeMin:  "192.137.0.33",
					RangeMax:  "192.137.0.62",
					Broadcast: "192.137.0.63",
				},
				{
					Name:      "IT",
					Mode:      Minimum,
					Size:      512,
					IP:        "192.137.0.64/23",
					Mask:      "255.255.254.0",
					RangeMin:  "192.137.0.65",
					RangeMax:  "192.137.2.62",
					Broadcast: "192.137.2.63",
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

func TestValidateModes(t *testing.T) {
	egs := []struct {
		input  Subnets
		output Subnets
	}{
		{
			input: Subnets{
				{Mode: Minimum},
				{Mode: Minimum},
			},
			output: Subnets{
				{Mode: Minimum},
				{Mode: Minimum},
			},
		},
		{
			input: Subnets{
				{Mode: Minimum},
				{Mode: Maximum},
				{Mode: Balanced},
			},
			output: Subnets{
				{Mode: Minimum},
				{Mode: Maximum},
				{Mode: Balanced},
			},
		},
		{
			input: Subnets{
				{Mode: Minimum},
				{Mode: Maximum},
				{Mode: Maximum},
				{Mode: Balanced},
			},
			output: Subnets{
				{Mode: Minimum},
				{Mode: Balanced},
				{Mode: Balanced},
				{Mode: Balanced},
			},
		},
	}

	for _, eg := range egs {
		got := validateModes(eg.input)
		if !reflect.DeepEqual(got, eg.output) {
			t.Errorf("validateModes(%v)\n exp: %v\n got: %v\n\n", eg.input, eg.output, got)
		}
	}
}
