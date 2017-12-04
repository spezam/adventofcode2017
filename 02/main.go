package main

// The spreadsheet consists of rows of apparently-random numbers. To make sure the
// recovery process is on the right track, they need you to calculate the
// spreadsheet's checksum. For each row, determine the difference between the
// largest value and the smallest value; the checksum is the sum of all of these
// differences.

// --- Part Two ---
// The goal is to find the only two numbers in each row where one evenly divides
// the other - that is, where the result of the division operation is a whole
// number. They would like you to find those numbers on each line, divide them,
// and add up each line's result.
//
// Answer part one: 32020
// Answer part two: 236

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

		z += ChecksumLine(s)
		k += EvenlyDivisible(s)
	}

	fmt.Println("Checksum is:", z)
	fmt.Println("Evenly divisible values sum is:", k)
}

func ChecksumLine(l []int) int {
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

func EvenlyDivisible(l []int) int {
	var s int

	for k, v := range l {
		for kk, vv := range l {
			// skip itself
			if k == kk {
				continue
			}

			// check remainder
			if v%vv == 0 {
				s += v / vv
			}
		}
	}

	return s
}

// eof
