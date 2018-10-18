package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}
	target := 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for index, val := range nums {
		sub := target - val
		if key, ok := m[sub]; ok {
			return []int{index, key}
		} else {
			m[val] = index
		}
	}
	return nil
}
