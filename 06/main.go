package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("day06")

	f, err := ioutil.ReadFile("input_test")
	if err != nil {
		panic(err)
	}

	mba := []int{}
	for _, v := range strings.Split(string(f), "\t") {
		vv, _ := strconv.Atoi(v)

		mba = append(mba, vv)
	}

	mbr := ring.New(len(strings.Split(string(f), "\t")))
	for _, v := range strings.Split(string(f), "\t") {
		vv, _ := strconv.Atoi(v)
		mbr.Value = vv

		mbr = mbr.Next()
	}

	// var seen = map[string]bool{}
	// var cycles int
	for range mba {
		maxIdx := maxValIndex(mbr)
		mbr = mbr.Move(maxIdx) // reset initial block
		max := mbr.Value.(int)
		mbr.Value = 0

		for i := 0; i < max; i++ {
			mbr = mbr.Next()
			fmt.Println(mbr.Value)
			mbr.Value = mbr.Value.(int) + 1
		}

		fmt.Println("--- ---")
		// mbr = mbr.Move(0)
		mbr.Do(func(x interface{}) {
			fmt.Println(x)
		})
		fmt.Println("--- ---")
		fmt.Println("---")
		// break

		// 	// seen[strings.Join(strconv.Itoa(mb))] = true
	}
}

func maxValIndex(r *ring.Ring) int {
	m := r.Value
	z, k := 0, 0

	r.Do(func(x interface{}) {
		if m.(int) < x.(int) {
			z = k // key
			m = x // value
		}
		k++
	})

	return z
}
