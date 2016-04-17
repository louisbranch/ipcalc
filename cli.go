package main

import (
	"bytes"
	"fmt"
	"strconv"
	"text/tabwriter"
)

func promptNetwork() (input Network, err error) {
	txt, err := prompt("IPv4 Network address (CIDR format)")
	if err != nil {
		return input, err
	}
	input.IP = txt

	txt, err = prompt("Number of subnets to create")
	if err != nil {
		return input, err
	}
	n, err := strconv.Atoi(txt)
	if err != nil {
		return input, fmt.Errorf("Error parsing number of subnets %s", err)
	}

	for i := 0; i < n; i++ {
		msg := fmt.Sprintf("Subnet #%d name", i+1)
		name, err := prompt(msg)
		if err != nil {
			return input, err
		}

		msg = fmt.Sprintf("Subnet #%d size", i+1)
		txt, err := prompt(msg)
		if err != nil {
			return input, err
		}
		size, err := strconv.Atoi(txt)
		if err != nil {
			return input, fmt.Errorf("Error parsing subnet size %s", err)
		}
		if n < 1 {
			return input, fmt.Errorf("Subnet size must be greater than 0")
		}

		msg = fmt.Sprintf("Subnet #%d mode (0 = Mininum, 1 = Maximum, 2 = Balanced)", i+1)
		txt, err = prompt(msg)
		if err != nil {
			return input, err
		}
		mode, err := strconv.Atoi(txt)
		if err != nil {
			return input, fmt.Errorf("Error parsing subnet mode %s", err)
		}
		if n < 0 || n > 3 {
			return input, fmt.Errorf("Invalid mode")
		}

		subnet := Subnet{Name: name, Size: size, Mode: mode}
		input.Subnets = append(input.Subnets, subnet)
	}

	return input, nil
}

func output(ip string, subnets Subnets) []byte {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 8, 8, 2, ' ', 0)

	fmt.Fprintf(w, "Network:\t%s\n", ip)
	fmt.Fprintf(w, "Subnets:\t%d\n", len(subnets))
	fmt.Fprintln(w, "---------\t")

	for _, subnet := range subnets {
		fmt.Fprintf(w, "Name:\t%s\n", subnet.Name)
		fmt.Fprintf(w, "Mode:\t%s\n", printMode(subnet.Mode))
		fmt.Fprintf(w, "Address:\t%s\n", subnet.IP)
		fmt.Fprintf(w, "Size:\t%d\n", subnet.Size-2)
		fmt.Fprintf(w, "Mask:\t%s\n", subnet.Mask)
		fmt.Fprintf(w, "Host Min:\t%s\n", subnet.RangeMin)
		fmt.Fprintf(w, "Host Max:\t%s\n", subnet.RangeMax)
		fmt.Fprintf(w, "Broadcast:\t%s\n", subnet.Broadcast)
		fmt.Fprintln(w, "---------\t")
	}

	w.Flush()
	return buf.Bytes()
}

func prompt(msg string) (string, error) {
	fmt.Printf("%s: ", msg)
	if !in.Scan() {
		return "", fmt.Errorf("Error parsing stdin %s", in.Err())
	}
	return in.Text(), nil
}

func printMode(mode int) string {
	switch mode {
	case 0:
		return "Mininum"
	case 1:
		return "Maximum"
	case 2:
		return "Balanced"
	default:
		return "Invalid"
	}
}
