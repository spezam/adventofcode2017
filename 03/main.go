package main

// Answer part one: 480
// Answer part two:

import (
	"fmt"
	"math"
)

func main() {
	var input float64 = 347991
	//var input float64 = 480
	fmt.Println("Input:", input)

	x, y := SpiralPoint(input)
	distance := math.Abs(x) + math.Abs(y)

	fmt.Println("Manhattan distance:", distance)
}

// function spiral(n)
//         k=ceil((sqrt(n)-1)/2)
//         t=2*k+1
//         m=t^2
//         t=t-1
//
//         if n>=m-t then return k-(m-n),-k        else m=m-t end
//         if n>=m-t then return -k,-k+(m-n)       else m=m-t end
//         if n>=m-t then return -k+(m-n),k else return k,k-(m-n-t) end
// end
func SpiralPoint(n float64) (float64, float64) {
	k := math.Ceil((math.Sqrt(n-1) / 2))
	var t float64 = 2*k + 1
	m := math.Pow(t, 2)
	t = t - 1

	if n >= m-t {
		return k - (m - n), -k
	} else {
		m = m - t

		if n >= m-t {
			return -k, -k + (m - n)
		} else {
			m = m - t

			if n >= m-t {
				return -k + (m - n), k
			} else {
				return k, k - (m - n - t)
			}
		}
	}
}
