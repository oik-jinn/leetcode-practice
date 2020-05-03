package main

import (
	"fmt"
	"sort"
	"os"
)

func main() {
	nums := []int{3, 2, 3}
	target := 6

	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum1(nums, target))
	os.Exit(1)
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

func twoSum1(nums []int, target int) []int {
	backs := append(nums[:0:0], nums...)
	sort.Ints(nums)

	start := 0
	end := len(nums)-1

	for start < end {
		if nums[start]+nums[end] > target {
			end--
		} else if(nums[start]+nums[end]<target) {
			start++
		} else if(nums[end]+nums[start]==target) {
			break
		}

	}

	result := make([]int, 2)
	index := 0
	one := nums[start]
	another := nums[end]
	for i, value := range backs {
		if one == value || another == value {
			result[index] = i
			index ++ 
		}
	}

	return result
}
