package main

import (
	"fmt"
	"os"
)

func main() {
	A := []int{5,7,7,8,8,10}
	A = []int{8}
	// A = []int{8,8}
	target := 8

	fmt.Println(searchRange(A, target))
	os.Exit(1)
}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums)-1

	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	// Find min index
	for left < right {
		mid := left + (right-left) / 2 
		if nums[mid] < target {
			left = mid +1
		} else {
			right = mid
		}
	}

	if nums[left] != target {
		return result
	}

	result[0] = left

	// Find max index
	right = len(nums) -1 
	for left < right {
		mid := left + (right-left) / 2 +1
		if nums[mid] > target {
			right = mid-1
		} else {
			left = mid	
		}
	}

	result[1] = right
	return result
}