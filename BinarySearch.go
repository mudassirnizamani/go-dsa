package main

import (
	"fmt"
	"math"
)

func binarySearch(haystack []int, needle int) bool {
	lo := 0
	hi := len(haystack)

	for {
		m := int(math.Floor(float64(lo + (hi-lo)/2)))
		v := haystack[m]

		if lo < hi {
			if v == needle {
				return true
			} else if v > needle {
				hi = int(m)
			} else {
				lo = m + 1
			}
		} else {
			return false
		}
	}
}

func main() {
	var foo []int = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	var test1 = binarySearch(foo, 3)

	fmt.Println(test1)
}
