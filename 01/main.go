package main

// --- Day 1: Inverse Captcha ---
// The captcha requires you to review a sequence of digits (your puzzle input)
// and find the sum of all digits that match the next digit in the list. The
// list is circular, so the digit after the last digit is the first digit in
// the list.
//
// --- Part Two ---
// Now, instead of considering the next digit, it wants you to consider the digit
// halfway around the circular list. That is, if your list contains 10 items, only
// include a digit in your sum if the digit 10/2 = 5 steps forward matches it.
//
// Answer part one: 1251
// Answer part two: 1244

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("captcha")
	if err != nil {
		panic(err)
	}

	c := string(f)
	z := Captcha(c)
	fmt.Println("Result for Captcha is:", z)

	z = CaptchaRing(c)
	fmt.Println("Result for Captcha (using ring datatype) is:", z)

	z = HalfCaptcha(c)
	fmt.Println("Result for Captcha (using ring type) is:", z)

}

// Captcha calc
func Captcha(c string) int {
	var s []int

	// to []int
	for _, v := range strings.Split(c, "") {
		i, _ := strconv.Atoi(v)
		s = append(s, i)
	}

	var z int
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			z += s[i]
		}
	}

	// circular list
	if s[0] == s[len(s)-1] {
		z += s[0]
	}

	return z
}

// CaptchaRing - Captcha calc using ring datatype
func CaptchaRing(c string) int {
	r := ring.New(len(c))
	for _, v := range strings.Split(c, "") {
		i, _ := strconv.Atoi(v)
		r.Value = i

		r = r.Next()
	}

	z := 0
	r.Do(func(x interface{}) {
		if x == r.Next().Value {
			z += x.(int)
		}

		// move cursor to next ring element
		r = r.Next()
	})

	return z
}

// HalfCaptcha ...
func HalfCaptcha(c string) int {
	sum := 0
	for i := 0; i < len(c)/2; i++ {
		first, _ := strconv.Atoi(string(c[i]))
		second, _ := strconv.Atoi(string(c[i+len(c)/2]))
		if first == second {
			sum += first * 2
		}
	}

	return sum
}

// eof
