package main

import (
	"fmt"
)

func main() {
	a := []int{1, 43, 5, 6, -1, 11, -2}
	//a := []int{1}
	target := 10
	fmt.Println(TwoNumberSum(a, target))
}

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
	var b []int

	for i := 0; i <= len(array)-1; i++ {
		for j := len(array) - 1; j >= 0; j-- {

			if array[i]+array[j] == target {
				if i == j {
					continue
				} else {
					b = append(b, array[i], array[j])
					return b
				}
			}
		}
	}
	return []int{}
}
