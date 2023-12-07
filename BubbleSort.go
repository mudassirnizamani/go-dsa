package main

import "fmt"

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				var tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func main() {
	var arr = []int{9, 3, 7, 4, 69, 420, 42}

	fmt.Println(bubbleSort(arr))
}
