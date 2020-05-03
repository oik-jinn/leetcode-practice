package main

import (
	"os"
	"fmt"
)

func main() {

	nums := [] int {-1, 0, 1, 2, -1, -4}
	fmt.Println(nums)
	os.Exit(1)

	result := threeSum(nums)
	fmt.Println(result)
	os.Exit(1)
}

func threeSum(nums []int) [][]int {
	
	result := make([][]int,2)
	result[0] = nums

	return result	
}