package main

import (
	"fmt"
	"math"
	"math/rand"
)

func twoCrystalBalls(breaks []bool) int {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jmpAmount
	fmt.Println("i: ", i)

	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	fmt.Println("i: ", i)
	i -= jmpAmount
	fmt.Println("i: ", i)
	for j := 0; j < jmpAmount && i < len(breaks); {
		fmt.Println(breaks[i])
		if breaks[i] {
			return i
		}
		i++
		j++
	}

	return -1
}

func makeSlice(idx int) []bool {
	data := make([]bool, 10000)
	for i := idx; i < 10000; i++ {
		data[int(i)] = true
	}
	return data
}

func main() {
	idx := int(math.Floor(float64(rand.Intn(10000))))
	data := makeSlice(idx)

	test1 := twoCrystalBalls(data)

	fmt.Println(test1)
}
