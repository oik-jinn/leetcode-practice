package main

import (
	"os"
	"fmt"
)

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	result := search(nums, target)
	fmt.Println(result)
	os.Exit(1)
}

func search(nums []int, target int) int {
	return target
}