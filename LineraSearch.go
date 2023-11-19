package main

import "fmt"

func linearSearch(haystack []int, needle int) bool {

	i := 0

	for i < len(haystack) {
		v := haystack[i]

		if v == needle {
			return true
		}

		i++
	}

	return false
}

func main() {
	var foo []int = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	var test1 = linearSearch(foo, 420)

	fmt.Println(test1)
}
