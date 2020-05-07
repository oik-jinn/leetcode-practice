package main

import (
	"fmt"
	"os"
)

func main() {
	A := []int{1,1,2}
	// A = []int{0,0,1,1,1,2,2,3,3,4}

	fmt.Println(removeDuplicates(A))
	os.Exit(1)
}

func removeDuplicates(nums []int) int {
	j := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}