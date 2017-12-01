package main

// The captcha requires you to review a sequence of digits (your puzzle input)
// and find the sum of all digits that match the next digit in the list. The
// list is circular, so the digit after the last digit is the first digit in
// the list.

import (
	"container/ring"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	c := flag.String("c", "", "Captcha")
	flag.Parse()

	if *c == "" {
		flag.PrintDefaults()
		return
	}

	z := Captcha(*c)
	fmt.Println("Result for Captcha", *c, "is:", z)
	z = CaptchaRing(*c)
	fmt.Println("Result for CaptchaRing", *c, "is:", z)
}

// Captcha calc
func Captcha(c string) int {
	s := []int{}

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

func CaptchaRing(c string) int {
	r := ring.New(len(c))
	for _, v := range strings.Split(c, "") {
		i, _ := strconv.Atoi(v)
		r.Value = i

		r = r.Next()
	}

	var z int
	r.Do(func(x interface{}) {
		if x == r.Next().Value {
			z += x.(int)
		}

		// move cursor to next ring element
		r = r.Next()
	})

	return z
}

// eof
