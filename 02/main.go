package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	defer f.Close()

	lines := csv.NewReader(f)
	lines.Comma = '\t'

	var k, z int
	for {
		r, err := lines.Read()
		if err == io.EOF {
			break
		}

		s := []int{}
		for _, v := range r {
			i, _ := strconv.Atoi(v)
			s = append(s, i)
		}

		z += checksumLine(s)
		k += evenlyDivisible(s)
	}

	fmt.Println("Checksum is:", z)
	fmt.Println("Evenly divisible values sum is:", k)
}

func checksumLine(l []int) int {
	s, b := l[0], l[0]

	for _, v := range l {
		// bigger
		if v > b {
			b = v
		}

		// smaller
		if v < s {
			s = v
		}
	}

	return b - s
}

func evenlyDivisible(l []int) int {
	var s int

	for k, v := range l {
		for i := 0; i < len(l); i++ {
			// skip itself
			if k != i {
				// check remainder
				if v%l[i] == 0 {
					s += v / l[i]
				}
			}

		}
	}

	return s
}

// eof
