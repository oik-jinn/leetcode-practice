package main

import (
	"fmt"
	"os"
)

func main() {
	A := []int{3,2,2,3}
	target := 3

	fmt.Println(removeElement(A, target))
	os.Exit(1)
}

func removeElement(nums []int, target int) int {
	j := 0
	for _, val := range nums {
		if val != target {
			nums[j] = val
			j++
		}
	}

	return j
}