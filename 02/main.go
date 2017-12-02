package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	var k, z int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := strings.Fields(scanner.Text())

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
			if k == i {
				continue
			}

			// check remainder
			if v%l[i] == 0 {
				s += v / l[i]
			}
		}
	}

	return s
}

// eof
